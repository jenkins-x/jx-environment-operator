package environments

import (
	"fmt"
	"sync/atomic"
	"time"

	"k8s.io/apimachinery/pkg/util/runtime"

	"github.com/jenkins-x/jx-logging/v3/pkg/log"

	v1 "github.com/jenkins-x/jx-api/v3/pkg/apis/jenkins.io/v1"

	"github.com/jenkins-x/jx-api/v3/pkg/client/clientset/versioned"
	informers "github.com/jenkins-x/jx-api/v3/pkg/client/informers/externalversions"
	"k8s.io/client-go/tools/cache"
)

type Informer struct {
	Client           versioned.Interface
	Namespace        string
	ResyncInterval   time.Duration
	informerFactory  informers.SharedInformerFactory
	activityInformer cache.SharedIndexInformer
	IsReady          *atomic.Value
	ClusterName      string
}

// Start starts the informers used to collect data and then send to the DB
func (i *Informer) Start() {

	i.informerFactory = informers.NewSharedInformerFactoryWithOptions(
		i.Client,
		i.ResyncInterval,
		informers.WithNamespace(i.Namespace),
	)

	stop := make(chan struct{})
	defer close(stop)

	// Kubernetes serves an utility to handle API crashes
	defer runtime.HandleCrash()

	i.activityInformer = i.informerFactory.Jenkins().V1().Environments().Informer()
	i.activityInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			environment := obj.(*v1.Environment)

			log.Logger().Infof("Added environment %s", environment.Name)
		},
	})

	// Starts all the shared informers that have been created by the factory so
	// far.
	i.informerFactory.Start(stop)
	// wait for the initial synchronization of the local cache.
	if !cache.WaitForCacheSync(stop, i.activityInformer.HasSynced) {
		runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
	i.IsReady.Store(true)
	<-stop
}
