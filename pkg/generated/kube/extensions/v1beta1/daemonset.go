// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	informers "k8s.io/client-go/informers/extensions/v1beta1"
	listers "k8s.io/client-go/listers/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type daemonSetInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.DaemonSetInformer = &daemonSetInformer{}

func NewDaemonSetInformer(f xnsinformers.SharedInformerFactory) informers.DaemonSetInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("daemonsets")
	informer := f.NamespacedResource(resource).Informer()

	return &daemonSetInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta1.DaemonSet{}),
	}
}

func (i *daemonSetInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *daemonSetInformer) Lister() listers.DaemonSetLister {
	return listers.NewDaemonSetLister(i.informer.GetIndexer())
}
