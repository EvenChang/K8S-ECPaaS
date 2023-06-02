// This is a generated file. Do not edit directly.
// Run hack/pin-dependency.sh to change pinned dependency versions.
// Run hack/update-vendor.sh to update go.mod files and the vendor directory.

module kubesphere.io/api

go 1.13

require (
	github.com/go-openapi/spec v0.19.7
	github.com/onsi/gomega v1.15.0
	github.com/projectcalico/libcalico-go v1.7.2-0.20191014160346-2382c6cdd056
	golang.org/x/net v0.0.0-20211209124913-491a49abca63
	istio.io/api v0.0.0-20201113182140-d4b7e3fc2b44
	k8s.io/api v0.21.11
	k8s.io/apiextensions-apiserver v0.21.11
	k8s.io/apimachinery v0.21.11
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20211110012726-3cc51fd1e909
	kubevirt.io/api v0.49.1
	kubevirt.io/containerized-data-importer-api v1.41.0 // indirect
	sigs.k8s.io/application v0.8.4-0.20201016185654-c8e2959e57a0
	sigs.k8s.io/controller-runtime v0.9.3
)

replace (
	github.com/go-logr/logr => github.com/go-logr/logr v0.4.0
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1
	github.com/openshift/client-go v0.0.0 => github.com/openshift/client-go v0.0.0-20210112165513-ebc401615f47
	k8s.io/client-go => k8s.io/client-go v0.21.2
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.8.0
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
	kubevirt.io/client-go => kubevirt.io/client-go v0.49.1
	sigs.k8s.io/application => sigs.k8s.io/application v0.8.4-0.20201016185654-c8e2959e57a0
)
