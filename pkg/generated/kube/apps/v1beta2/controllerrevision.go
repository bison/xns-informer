// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta2

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta2 "k8s.io/api/apps/v1beta2"
	informers "k8s.io/client-go/informers/apps/v1beta2"
	listers "k8s.io/client-go/listers/apps/v1beta2"
	"k8s.io/client-go/tools/cache"
)

type controllerRevisionInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.ControllerRevisionInformer = &controllerRevisionInformer{}

func NewControllerRevisionInformer(f xnsinformers.SharedInformerFactory) informers.ControllerRevisionInformer {
	resource := v1beta2.SchemeGroupVersion.WithResource("controllerrevisions")
	informer := f.NamespacedResource(resource).Informer()

	return &controllerRevisionInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta2.ControllerRevision{}),
	}
}

func (i *controllerRevisionInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *controllerRevisionInformer) Lister() listers.ControllerRevisionLister {
	return listers.NewControllerRevisionLister(i.informer.GetIndexer())
}
