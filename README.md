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

fix-bug proto: duplicate proto type registered

```
#       modified:   Godeps/Godeps.json
#       modified:   Godeps/LICENSES
#       modified:   vendor/BUILD
#       modified:   vendor/github.com/libopenstorage/openstorage/api/BUILD
#       modified:   vendor/github.com/libopenstorage/openstorage/api/api.pb.go
$cd vendor/github.com/libopenstorage/openstorage/
$protoc --go_out=. ./api/api.proto
```


#add labels to containers

```
kind: Deployment
metadata:
  name: tlabels
  namespace: default
spec:
  replicas: 1
  template:
    metadata:
      labels:
        run: tlabels
      annotations:
        froad-containers-labels/aliyun.logs.standard_srv_operating-report_debug: /data/logs/operating-report_34657_debug
        froad-containers-labels/aliyun.logs.standard_srv_operating-report_error: /data/logs/operating-report_34657_error
    spec:
      containers:
      - name: tlabels
        image: index.tenxcloud.com/docker_library/nginx
        ports:
        - containerPort: 80
      userDefineNet: vlan608
```
