// Code generated by xns-informer-gen. DO NOT EDIT.

package v2beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/autoscaling/v2beta1"
)

type Interface interface {
	HorizontalPodAutoscalers() informers.HorizontalPodAutoscalerInformer
}

type version struct {
	factory xnsinformers.SharedInformerFactory
}

func New(factory xnsinformers.SharedInformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) HorizontalPodAutoscalers() informers.HorizontalPodAutoscalerInformer {
	return &horizontalPodAutoscalerInformer{factory: v.factory}
}