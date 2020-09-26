// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/storage/v1beta1"
	informers "k8s.io/client-go/informers/storage/v1beta1"
	listers "k8s.io/client-go/listers/storage/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type cSINodeInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.CSINodeInformer = &cSINodeInformer{}

func NewCSINodeInformer(f xnsinformers.SharedInformerFactory) informers.CSINodeInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("csinodes")
	informer := f.ClusterResource(resource).Informer()

	return &cSINodeInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta1.CSINode{}),
	}
}

func (i *cSINodeInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *cSINodeInformer) Lister() listers.CSINodeLister {
	return listers.NewCSINodeLister(i.informer.GetIndexer())
}
