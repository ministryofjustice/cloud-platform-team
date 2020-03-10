package util

import (
	api_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func GetPodsSharedIndexInformer(client kubernetes.Interface) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		//the ListWatch contains two different functions that our
		//informer requires: ListFunc to take care of listing and watching
		//the resources we want to handle.
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				//list all of the pods (core resource) in the default namespace
				return client.CoreV1().Pods(meta_v1.NamespaceDefault).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				//watch all of the pods (core resource) in the default namespace
				return client.CoreV1().Pods(meta_v1.NamespaceDefault).Watch(options)
			},
		},
		&api_v1.Pod{}, //the target type (Pod)
		0,			   // no resync (period of 0)
		cache.Indexers{},
	)
}

func CreateWorkingQueue() workqueue.RateLimitingInterface {
	//a result of listing or watching, we can add identifying key to the queue
	//so that it can be handled in the handler
	return workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
}
