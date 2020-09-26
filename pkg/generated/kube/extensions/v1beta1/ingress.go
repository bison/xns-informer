// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	informers "k8s.io/client-go/informers/extensions/v1beta1"
	listers "k8s.io/client-go/listers/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type ingressInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.IngressInformer = &ingressInformer{}

func NewIngressInformer(f xnsinformers.SharedInformerFactory) informers.IngressInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("ingresses")
	informer := f.NamespacedResource(resource).Informer()

	return &ingressInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta1.Ingress{}),
	}
}

func (i *ingressInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *ingressInformer) Lister() listers.IngressLister {
	return listers.NewIngressLister(i.informer.GetIndexer())
}
