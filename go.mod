module github.com/jenkins-x/jx-environment-operator

go 1.15

require (
	github.com/Azure/go-autorest/autorest/adal v0.9.11 // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.17+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/google/go-github/v29 v29.0.3 // indirect
	github.com/googleapis/gnostic v0.5.4 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/jenkins-x/jx-api/v3 v3.0.3
	github.com/jenkins-x/jx-api/v4 v4.0.24 // indirect
	github.com/jenkins-x/jx-gitops v0.0.548 // indirect
	github.com/jenkins-x/jx-kube-client/v3 v3.0.2
	github.com/jenkins-x/jx-logging/v3 v3.0.3
	github.com/jenkins-x/jx-preview v0.0.155
	github.com/jenkins-x/jx-secret v0.0.226 // indirect
	github.com/jenkins-x/lighthouse v0.0.923 // indirect
	github.com/jenkins-x/structs v1.1.0 // indirect
	github.com/klauspost/cpuid v1.2.2 // indirect
	github.com/knative/build v0.1.2 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/nats-io/gnatsd v1.4.1 // indirect
	github.com/nats-io/go-nats v1.7.0 // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/test-infra v0.0.0-20200630141629-15f40fe97047 // indirect
	sigs.k8s.io/structured-merge-diff v1.0.1 // indirect
	sigs.k8s.io/testing_frameworks v0.1.1 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.19.2
