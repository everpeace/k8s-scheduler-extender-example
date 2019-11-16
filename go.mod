module k8s-scheduler-extender-example

go 1.13

require (
	github.com/comail/colog v0.0.0-20160416085026-fba8e7b1f46c
	github.com/julienschmidt/httprouter v1.1.0
)

replace (
	k8s.io/api => ./kubernetes/staging/src/k8s.io/api
	k8s.io/apiextensions-apiserver => ./kubernetes/staging/src/k8s.io/apiextensions-apiserver
	k8s.io/apimachinery => ./kubernetes/staging/src/k8s.io/apimachinery
	k8s.io/apiserver => ./kubernetes/staging/src/k8s.io/apiserver
	k8s.io/cli-runtime => ./kubernetes/staging/src/k8s.io/cli-runtime
	k8s.io/client-go => ./kubernetes/staging/src/k8s.io/client-go
	k8s.io/cloud-provider => ./kubernetes/staging/src/k8s.io/cloud-provider
	k8s.io/cluster-bootstrap => ./kubernetes/staging/src/k8s.io/cluster-bootstrap
	k8s.io/code-generator => ./kubernetes/staging/src/k8s.io/code-generator
	k8s.io/component-base => ./kubernetes/staging/src/k8s.io/component-base
	k8s.io/cri-api => ./kubernetes/staging/src/k8s.io/cri-api
	k8s.io/csi-translation-lib => ./kubernetes/staging/src/k8s.io/csi-translation-lib
	k8s.io/gengo => k8s.io/gengo v0.0.0-20190116091435-f8a0810f38af
	k8s.io/heapster => k8s.io/heapster v1.2.0-beta.1
	k8s.io/klog => k8s.io/klog v0.3.1
	k8s.io/kube-aggregator => ./kubernetes/staging/src/k8s.io/kube-aggregator
	k8s.io/kube-controller-manager => ./kubernetes/staging/src/k8s.io/kube-controller-manager
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190816220812-743ec37842bf
	k8s.io/kube-proxy => ./kubernetes/staging/src/k8s.io/kube-proxy
	k8s.io/kube-scheduler => ./kubernetes/staging/src/k8s.io/kube-scheduler
	k8s.io/kubectl => ./kubernetes/staging/src/k8s.io/kubectl
	k8s.io/kubelet => ./kubernetes/staging/src/k8s.io/kubelet
	k8s.io/kubernetes => ./kubernetes
	k8s.io/legacy-cloud-providers => ./kubernetes/staging/src/k8s.io/legacy-cloud-providers
	k8s.io/metrics => ./kubernetes/staging/src/k8s.io/metrics
	k8s.io/node-api => ./kubernetes/staging/src/k8s.io/node-api
	k8s.io/repo-infra => k8s.io/repo-infra v0.0.0-20181204233714-00fe14e3d1a3
	k8s.io/sample-apiserver => ./kubernetes/staging/src/k8s.io/sample-apiserver
	k8s.io/sample-cli-plugin => ./kubernetes/staging/src/k8s.io/sample-cli-plugin
	k8s.io/sample-controller => ./kubernetes/staging/src/k8s.io/sample-controller
)
