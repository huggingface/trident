// Copyright 2016 NetApp, Inc. All Rights Reserved.

package main

import (
	"reflect"
	"strings"
	"testing"

	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/pkg/version"

	tridentrest "github.com/netapp/trident/frontend/rest"
	"github.com/netapp/trident/k8s_client"
	"github.com/netapp/trident/storage"
)

func TestValidKubeVersion(t *testing.T) {
	k8sVersion := &version.Info{
		Major: "1",
		Minor: "5",
	}
	launcher := &Launcher{}
	supported, err := launcher.ValidateVersion(k8sVersion)
	if !supported || err != nil {
		t.Fatalf("The test for a valid Kubernetes version failed: %s", err)
	}
}

func TestInvalidKubeVersion(t *testing.T) {
	k8sVersion := &version.Info{
		Major: "1",
		Minor: "3",
	}
	launcher := &Launcher{}
	supported, err := launcher.ValidateVersion(k8sVersion)
	if supported || err != nil {
		t.Fatalf("The test for an invalid Kubernetes version failed: %s", err)
	}
}

func TestValidOpenShiftVersion(t *testing.T) {
	osVersion := &version.Info{
		Major: "3",
		Minor: "5",
	}
	launcher := &Launcher{}
	supported, err := launcher.ValidateVersion(osVersion)
	if !supported || err != nil {
		t.Fatalf("The test for a valid OpenShift version failed: %s", err)
	}
}

func TestInvalidOpenShiftVersion(t *testing.T) {
	osVersion := &version.Info{
		Major: "3",
		Minor: "3",
	}
	launcher := &Launcher{}
	supported, err := launcher.ValidateVersion(osVersion)
	if supported || err != nil {
		t.Fatalf("The test for an invalid OpenShift version failed: %s", err)
	}
}

func TestTridentClientVolume(t *testing.T) {
	tridentClientFailMatrix := map[string]bool{}
	tridentClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	getVolResponse, err := tridentClient.GetVolume("trident")
	if err != nil || getVolResponse.Error != "Volume wasn't found" {
		t.Fatal("Error in retrieving a non-existent volume!")
	}
	_, err =
		tridentClient.AddVolume(&storage.VolumeConfig{Name: "trident"})
	if err != nil {
		t.Fatal("Error in creating a volume!")
	}
	getVolResponse, err = tridentClient.GetVolume("trident")
	if err != nil || getVolResponse.Error != "" ||
		getVolResponse.Volume.Config.Name != "trident" {
		t.Fatal("Error in retrieving a nonexistent volume!")
	}
	deleteResponse, err := tridentClient.DeleteVolume("trident")
	if err != nil || deleteResponse.Error != "" {
		t.Fatal("Error in deleting a volume!")
	}
	deleteResponse, err = tridentClient.DeleteVolume("trident")
	if err != nil || deleteResponse.Error == "" {
		t.Fatal("Deleting a volume should succeed only once!")
	}
}

func TestKubeSnapshotStateValid(t *testing.T) {
	kubeClientFailMatrix := make(map[string]bool, 0)
	kubeClient := k8s_client.NewFakeKubeClient(kubeClientFailMatrix)
	snapshotBefore := kubeClient.SnapshotState()
	snapshotAfter := kubeClient.SnapshotState()
	if !reflect.DeepEqual(snapshotBefore, snapshotAfter) {
		t.Fatal("Kubernetes state shouldn't have changed!")
	}
}

func TestKubeSnapshotStateInvalid(t *testing.T) {
	var err error
	kubeClientFailMatrix := make(map[string]bool, 0)
	kubeClient := k8s_client.NewFakeKubeClient(kubeClientFailMatrix)
	snapshotBefore := kubeClient.SnapshotState()

	tridentDeployment := &v1beta1.Deployment{}
	tridentDeployment.Name = "trident"
	tridentDeployment, err = kubeClient.CreateDeployment(tridentDeployment)
	if err != nil {
		t.Fatal(err)
	}

	snapshotAfter := kubeClient.SnapshotState()
	if reflect.DeepEqual(snapshotBefore, snapshotAfter) {
		t.Fatal("Kubernetes state should have changed!")
	}
}

func TestExistingDeployment(t *testing.T) {
	var err error
	// Creating the parameters for launcher
	tridentClientFailMatrix := map[string]bool{}
	kubeClientFailMatrix := make(map[string]bool, 0)
	kubeClient := k8s_client.NewFakeKubeClient(kubeClientFailMatrix)
	tridentClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	tridentEphemeralClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	tridentDeployment := &v1beta1.Deployment{}
	tridentDeployment.Name = "trident"

	// Configuring the environment
	tridentDeployment, err = kubeClient.CreateDeployment(tridentDeployment)
	if err != nil {
		t.Fatal(err)
	}

	// Get the state of the Kubernetes cluster before running launcher
	snapshotBefore := kubeClient.SnapshotState()

	// Running launcher
	launcher := NewLauncher(kubeClient, tridentClient, tridentEphemeralClient,
		tridentDeployment)
	errors := launcher.Run()
	if len(errors) != 1 {
		t.Fatal("Launcher should have failed with the preexisting deployment!")
	}
	if !strings.Contains(errors[0].Error(),
		"Launcher detected a preexisting deployment") {
		t.Fatal("Launcher returned an incorrect error!")
	}

	// Make sure launcher didn't change the state of the Kubernetes cluster
	snapshotAfter := kubeClient.SnapshotState()
	if !reflect.DeepEqual(snapshotBefore, snapshotAfter) {
		t.Fatal("Launcher didn't clean up state properly!")
	}
}

func TestExistingDeploymentFailure(t *testing.T) {
	var err error
	// Creating the parameters for launcher
	tridentClientFailMatrix := map[string]bool{}
	kubeClientFailMatrix := map[string]bool{
		"GetDeployment": true,
	}
	kubeClient := k8s_client.NewFakeKubeClient(kubeClientFailMatrix)
	tridentClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	tridentEphemeralClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	tridentDeployment := &v1beta1.Deployment{}
	tridentDeployment.Name = "trident"

	// Configuring the environment
	tridentDeployment, err = kubeClient.CreateDeployment(tridentDeployment)
	if err != nil {
		t.Fatal(err)
	}

	// Get the state of the Kubernetes cluster before running launcher
	snapshotBefore := kubeClient.SnapshotState()

	// Running launcher
	launcher := NewLauncher(kubeClient, tridentClient, tridentEphemeralClient,
		tridentDeployment)
	errors := launcher.Run()
	if len(errors) != 1 {
		t.Fatal("Launcher should have failed!")
	}
	if !strings.Contains(errors[0].Error(),
		"Launcher couldn't establish the presence of deployment") {
		t.Fatal("Launcher returned an incorrect error!")
	}

	// Make sure launcher didn't change the state of the Kubernetes cluster
	snapshotAfter := kubeClient.SnapshotState()
	if !reflect.DeepEqual(snapshotBefore, snapshotAfter) {
		t.Fatal("Launcher didn't clean up state properly!")
	}
}

func TestExistingPVCFailure(t *testing.T) {
	var err error
	// Creating the parameters for launcher
	tridentClientFailMatrix := map[string]bool{}
	kubeClientFailMatrix := map[string]bool{
		"GetPVC": true,
	}
	kubeClient := k8s_client.NewFakeKubeClient(kubeClientFailMatrix)
	tridentClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	tridentEphemeralClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	tridentDeployment := &v1beta1.Deployment{}
	tridentDeployment.Name = "trident"

	// Configuring the environment
	tridentPVC := &v1.PersistentVolumeClaim{}
	tridentPVC.Name = "trident"
	tridentPVC, err = kubeClient.CreatePVC(tridentPVC)
	if err != nil {
		t.Fatal(err)
	}

	// Get the state of the Kubernetes cluster before running launcher
	snapshotBefore := kubeClient.SnapshotState()

	// Running launcher
	launcher := NewLauncher(kubeClient, tridentClient, tridentEphemeralClient,
		tridentDeployment)
	errors := launcher.Run()
	if len(errors) != 1 {
		t.Fatal("Launcher should have failed!")
	}
	if !strings.Contains(errors[0].Error(),
		"Launcher couldn't establish the presence of PVC") {
		t.Fatal("Launcher returned an incorrect error!")
	}

	// Make sure launcher didn't change the state of the Kubernetes cluster
	snapshotAfter := kubeClient.SnapshotState()
	if !reflect.DeepEqual(snapshotBefore, snapshotAfter) {
		t.Fatal("Launcher didn't clean up state properly!")
	}
}

func TestPrexistingBoundPVCFailedDeployment(t *testing.T) {
	var err error
	// Creating the parameters for launcher
	tridentClientFailMatrix := map[string]bool{}
	kubeClientFailMatrix := map[string]bool{
		"CreateDeployment": true,
	}
	kubeClient := k8s_client.NewFakeKubeClient(kubeClientFailMatrix)
	tridentClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	tridentEphemeralClient := tridentrest.NewFakeTridentClient(tridentClientFailMatrix)
	tridentDeployment := &v1beta1.Deployment{}
	tridentDeployment.Name = "trident"

	// Configuring the environment
	tridentPVC := &v1.PersistentVolumeClaim{}
	tridentPVC.Name = "trident"
	tridentPVC.Status.Phase = v1.ClaimBound
	tridentPVC, err = kubeClient.CreatePVC(tridentPVC)
	if err != nil {
		t.Fatal(err)
	}

	// Get the state of the Kubernetes cluster before running launcher
	snapshotBefore := kubeClient.SnapshotState()

	// Running launcher
	launcher := NewLauncher(kubeClient, tridentClient, tridentEphemeralClient,
		tridentDeployment)
	errors := launcher.Run()
	if len(errors) != 1 {
		t.Fatal("Launcher should have failed!")
	}
	if !strings.Contains(errors[0].Error(),
		"CreateDeployment failed") {
		t.Fatal("Launcher returned an incorrect error!")
	}

	// Make sure launcher didn't change the state of the Kubernetes cluster
	snapshotAfter := kubeClient.SnapshotState()
	if !reflect.DeepEqual(snapshotBefore, snapshotAfter) {
		t.Fatal("Launcher didn't clean up state properly!")
	}
}
