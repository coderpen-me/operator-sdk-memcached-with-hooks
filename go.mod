module github.com/example/memcached-operator

go 1.16

require (
	cloud.google.com/go v0.72.0 // indirect
	github.com/Kangaroux/go-map-schema v0.6.1
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5 // indirect
	k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	sigs.k8s.io/controller-runtime v0.9.2
)

replace github.com/kong/kubernetes-ingress-controller => github.com/kong/kubernetes-ingress-controller v1.2.0
