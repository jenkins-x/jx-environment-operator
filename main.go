package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jenkins-x/jx-environment-operator/pkg/handlers"
	"github.com/jenkins-x/jx-environment-operator/pkg/previews"
	"github.com/jenkins-x/jx-environment-operator/pkg/secrets"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/jenkins-x/jx-environment-operator/pkg/environments"

	"github.com/jenkins-x/jx-kube-client/v3/pkg/kubeclient"
	"github.com/jenkins-x/jx-api/v3/pkg/client/clientset/versioned"
	pversioned "github.com/jenkins-x/jx-preview/pkg/client/clientset/versioned"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
)

var (
	options struct {
		jxClient       versioned.Interface
		previewClient       pversioned.Interface
		kClient       kubernetes.Interface
		namespace      string
		resyncInterval time.Duration
		printVersion   bool
		port           string
		clusterName    string
	}

	// these are set at compile time by GoReleaser through LD Flags
	version  = "dev"
	revision = "unknown"
	date     = "now"
)

func init() {
	flag.StringVar(&options.namespace, "namespace", "jx", "namespace to watch resources")
	flag.DurationVar(&options.resyncInterval, "resync-interval", 1*time.Minute, "resync interval between full re-list operations")
	flag.BoolVar(&options.printVersion, "version", false, "Print the version")
	flag.StringVar(&options.port, "port", "8080", "port the health endpoint should listen on")

	err := setClients()
	if err != nil {
		log.Logger().Fatalf("failed to validate options: %v", err)
	}

	// can be empty string
	options.clusterName = os.Getenv("CLUSTER_NAME")
}

func main() {

	flag.Parse()

	if options.printVersion {
		fmt.Printf("Version %s - Revision %s - Date %s", version, revision, date)
		return
	}

	log.Logger().Infof("starting")
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		(&environments.Informer{
			Client:         options.jxClient,
			Namespace:      options.namespace,
			ResyncInterval: options.resyncInterval,
			IsReady:        isReady,
			ClusterName:    options.clusterName,
		}).Start()
	}()

	go func() {
		(&previews.Informer{
			Client:         options.previewClient,
			Namespace:      options.namespace,
			ResyncInterval: options.resyncInterval,
			IsReady:        isReady,
			ClusterName:    options.clusterName,
		}).Start()
	}()

	go func() {
		(&secrets.Informer{
			Client:         options.kClient,
			Namespace:      options.namespace,
			ResyncInterval: options.resyncInterval,
			IsReady:        isReady,
			ClusterName:    options.clusterName,
		}).Start()
	}()
	startHealthEndpoint(isReady)

}

func startHealthEndpoint(isReady *atomic.Value) {
	r := handlers.Router(isReady)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", options.port),
		Handler: r,
	}
	go func() {
		log.Logger().Fatal(srv.ListenAndServe())
	}()
	log.Logger().Infof("The service is ready to listen and serve.")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Logger().Infof("Got SIGINT...")
	case syscall.SIGTERM:
		log.Logger().Infof("Got SIGTERM...")
	}

	log.Logger().Infof("The service is shutting down...")
	err := srv.Shutdown(context.Background())
	if err != nil {
		log.Logger().Fatalf("failed to shutdown cleanly %v", err)
	}
	log.Logger().Infof("Done")
}

func setClients() error {

	f := kubeclient.NewFactory()
	cfg, err := f.CreateKubeConfig()
	if err != nil {
		return errors.Wrapf(err, "failed to get kubernetes config")
	}

	options.jxClient, err = versioned.NewForConfig(cfg)
	if err != nil {
		return errors.Wrapf(err, "error building jx client")
	}

	options.previewClient, err = pversioned.NewForConfig(cfg)
	if err != nil {
		return errors.Wrapf(err, "error building jx client")
	}

	options.kClient, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		return errors.Wrapf(err, "error building jx client")
	}
	return nil
}