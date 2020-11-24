module github.com/jenkins-x/jx-environment-operator

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/jenkins-x/jx-api/v3 v3.0.3
	github.com/jenkins-x/jx-kube-client/v3 v3.0.1
	github.com/jenkins-x/jx-logging/v3 v3.0.2
	github.com/jenkins-x/jx-preview v0.0.122
	github.com/jenkins-x/structs v1.1.0 // indirect
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.19.2
