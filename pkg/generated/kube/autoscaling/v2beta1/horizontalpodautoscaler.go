// Code generated by xns-informer-gen. DO NOT EDIT.

package v2beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	informers "k8s.io/client-go/informers/autoscaling/v2beta1"
	listers "k8s.io/client-go/listers/autoscaling/v2beta1"
	"k8s.io/client-go/tools/cache"
)

type horizontalPodAutoscalerInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.HorizontalPodAutoscalerInformer = &horizontalPodAutoscalerInformer{}

func NewHorizontalPodAutoscalerInformer(f xnsinformers.SharedInformerFactory) informers.HorizontalPodAutoscalerInformer {
	resource := v2beta1.SchemeGroupVersion.WithResource("horizontalpodautoscalers")
	informer := f.NamespacedResource(resource).Informer()

	return &horizontalPodAutoscalerInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v2beta1.HorizontalPodAutoscaler{}),
	}
}

func (i *horizontalPodAutoscalerInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *horizontalPodAutoscalerInformer) Lister() listers.HorizontalPodAutoscalerLister {
	return listers.NewHorizontalPodAutoscalerLister(i.informer.GetIndexer())
}
