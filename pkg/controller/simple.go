package controller

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ministryofjustice/cloud-platform-team-operator"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

//controller struct defines how a controller should encapsulate
//logging, client connectivity, informing (list and watching)
//queueing, and handling of resource changes.
type Controller struct {
	Logger   *log.Entry
	Clienset kubernetes.Interface
	Queue    workqueue.RateLimitingInterface
	Informer cache.SharedIndexInformer
	Handler  handler.SimpleHandler
}
