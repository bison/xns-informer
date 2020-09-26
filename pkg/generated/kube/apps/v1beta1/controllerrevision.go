// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/apps/v1beta1"
	informers "k8s.io/client-go/informers/apps/v1beta1"
	listers "k8s.io/client-go/listers/apps/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type controllerRevisionInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.ControllerRevisionInformer = &controllerRevisionInformer{}

func NewControllerRevisionInformer(f xnsinformers.SharedInformerFactory) informers.ControllerRevisionInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("controllerrevisions")
	informer := f.NamespacedResource(resource).Informer()

	return &controllerRevisionInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta1.ControllerRevision{}),
	}
}

func (i *controllerRevisionInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *controllerRevisionInformer) Lister() listers.ControllerRevisionLister {
	return listers.NewControllerRevisionLister(i.informer.GetIndexer())
}
