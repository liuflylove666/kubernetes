# Kubernetes
----
##### You have a working [Go environment].

```
$ go get -d k8s.io/kubernetes
$ cd $GOPATH/src/k8s.io/kubernetes
$hack/update-codecgen.sh
$hack/update-generated-runtime.sh
$hack/update-generated-protobuf.sh
$ make
or
$cd ./release/rpm  
$./docker-build.sh 
```
####add userDefineNet-POD###
add userDefineNet

```
modify files:
#       modified:   README.md
#       modified:   pkg/api/types.go
#       modified:   pkg/api/v1/generated.proto
#       modified:   pkg/api/v1/types.go
#       modified:   pkg/api/v1/types_swagger_doc_generated.go
#       modified:   pkg/kubelet/apis/cri/v1alpha1/runtime/api.proto
#       modified:   pkg/kubelet/dockershim/docker_sandbox.go
#       modified:   pkg/kubelet/kuberuntime/kuberuntime_sandbox.go
#       modified:   staging/src/k8s.io/client-go/pkg/api/types.go
#       modified:   staging/src/k8s.io/client-go/pkg/api/v1/types.go
#       modified:   staging/src/k8s.io/client-go/pkg/api/v1/types_swagger_doc_generated.go
#       modified:   vendor/github.com/docker/engine-api/types/configs.go
```
