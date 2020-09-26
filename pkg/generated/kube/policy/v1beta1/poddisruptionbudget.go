// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/api/policy/v1beta1"
	informers "k8s.io/client-go/informers/policy/v1beta1"
	listers "k8s.io/client-go/listers/policy/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type podDisruptionBudgetInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.PodDisruptionBudgetInformer = &podDisruptionBudgetInformer{}

func NewPodDisruptionBudgetInformer(f xnsinformers.SharedInformerFactory) informers.PodDisruptionBudgetInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("poddisruptionbudgets")
	informer := f.NamespacedResource(resource).Informer()

	return &podDisruptionBudgetInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1beta1.PodDisruptionBudget{}),
	}
}

func (i *podDisruptionBudgetInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *podDisruptionBudgetInformer) Lister() listers.PodDisruptionBudgetLister {
	return listers.NewPodDisruptionBudgetLister(i.informer.GetIndexer())
}
