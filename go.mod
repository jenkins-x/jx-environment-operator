module github.com/jenkins-x/jx-environment-operator

go 1.15

require (
	github.com/Azure/go-autorest/autorest/adal v0.9.10 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/googleapis/gnostic v0.4.2 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/jenkins-x/jx-api/v3 v3.0.3
	github.com/jenkins-x/jx-gitops v0.0.528 // indirect
	github.com/jenkins-x/jx-kube-client/v3 v3.0.1
	github.com/jenkins-x/jx-logging/v3 v3.0.3
	github.com/jenkins-x/jx-preview v0.0.142
	github.com/jenkins-x/jx-secret v0.0.209 // indirect
	github.com/jenkins-x/lighthouse v0.0.908 // indirect
	github.com/jenkins-x/structs v1.1.0 // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.19.2
