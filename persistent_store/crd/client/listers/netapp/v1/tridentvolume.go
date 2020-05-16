// Copyright 2020 NetApp, Inc. All Rights Reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TridentVolumeLister helps list TridentVolumes.
type TridentVolumeLister interface {
	// List lists all TridentVolumes in the indexer.
	List(selector labels.Selector) (ret []*v1.TridentVolume, err error)
	// TridentVolumes returns an object that can list and get TridentVolumes.
	TridentVolumes(namespace string) TridentVolumeNamespaceLister
	TridentVolumeListerExpansion
}

// tridentVolumeLister implements the TridentVolumeLister interface.
type tridentVolumeLister struct {
	indexer cache.Indexer
}

// NewTridentVolumeLister returns a new TridentVolumeLister.
func NewTridentVolumeLister(indexer cache.Indexer) TridentVolumeLister {
	return &tridentVolumeLister{indexer: indexer}
}

// List lists all TridentVolumes in the indexer.
func (s *tridentVolumeLister) List(selector labels.Selector) (ret []*v1.TridentVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentVolume))
	})
	return ret, err
}

// TridentVolumes returns an object that can list and get TridentVolumes.
func (s *tridentVolumeLister) TridentVolumes(namespace string) TridentVolumeNamespaceLister {
	return tridentVolumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TridentVolumeNamespaceLister helps list and get TridentVolumes.
type TridentVolumeNamespaceLister interface {
	// List lists all TridentVolumes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TridentVolume, err error)
	// Get retrieves the TridentVolume from the indexer for a given namespace and name.
	Get(name string) (*v1.TridentVolume, error)
	TridentVolumeNamespaceListerExpansion
}

// tridentVolumeNamespaceLister implements the TridentVolumeNamespaceLister
// interface.
type tridentVolumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TridentVolumes in the indexer for a given namespace.
func (s tridentVolumeNamespaceLister) List(selector labels.Selector) (ret []*v1.TridentVolume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentVolume))
	})
	return ret, err
}

// Get retrieves the TridentVolume from the indexer for a given namespace and name.
func (s tridentVolumeNamespaceLister) Get(name string) (*v1.TridentVolume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tridentvolume"), name)
	}
	return obj.(*v1.TridentVolume), nil
}
