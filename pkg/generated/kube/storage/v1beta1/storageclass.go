// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/storage/v1beta1"
	informers "k8s.io/client-go/informers/storage/v1beta1"
	listers "k8s.io/client-go/listers/storage/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type storageClassInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.StorageClassInformer = &storageClassInformer{}

func NewStorageClassInformer(f xnsinformers.SharedInformerFactory) informers.StorageClassInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("storageclasses")
	informer := f.ClusterResource(resource).Informer()

	return &storageClassInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta1.StorageClass{}),
	}
}

func (i *storageClassInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *storageClassInformer) Lister() listers.StorageClassLister {
	return listers.NewStorageClassLister(i.informer.GetIndexer())
}
