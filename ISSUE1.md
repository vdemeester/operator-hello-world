# Error when building as arbitrary user

```shell
$ operator-sdk build mariolet/podset-operator
go: finding github.com/spf13/pflag v1.0.5
go: finding sigs.k8s.io/controller-runtime v0.4.0
go: finding github.com/operator-framework/operator-sdk v0.12.0
go: finding k8s.io/api latest
go: finding k8s.io/apimachinery latest
go: finding k8s.io/client-go v11.0.0+incompatible
go: writing stat cache: open /go/pkg/mod/cache/download/sigs.k8s.io/controller-runtime/@v/v0.4.0.info317796775.tmp: permission denied
go: downloading sigs.k8s.io/controller-runtime v0.4.0
go: writing stat cache: open /go/pkg/mod/cache/download/github.com/spf13/pflag/@v/v1.0.5.info848108701.tmp: permission denied
go: downloading github.com/spf13/pflag v1.0.5
go: writing stat cache: open /go/pkg/mod/cache/download/github.com/operator-framework/operator-sdk/@v/v0.12.0.info717874195.tmp: permission denied
go: writing stat cache: open /go/pkg/mod/cache/download/k8s.io/client-go/@v/v11.0.0+incompatible.info321743539.tmp: permission denied
go: finding github.com/go-openapi/spec v0.19.4
go: finding k8s.io/kube-openapi latest
go: writing stat cache: open /go/pkg/mod/cache/download/github.com/go-openapi/spec/@v/v0.19.4.info45436501.tmp: permission denied
go: downloading github.com/go-openapi/spec v0.19.4
go: downloading k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
build github.com/l0rd/operator-hello-world/cmd/manager: cannot load github.com/go-openapi/spec: open /go/pkg/mod/cache/download/github.com/go-openapi/spec/@v/v0.19.4.lock: permission denied
Error: failed to build operator binary: (failed to exec []string{"go", "build", "-o", "/projects/src/github.com/l0rd/operator-hello-world/build/_output/bin/operator-hello-world", "-gcflags", "all=-trimpath=/projects/src/github.com/l0rd", "-asmflags", "all=-trimpath=/projects/src/github.com/l0rd", "github.com/l0rd/operator-hello-world/cmd/manager"}: exit status 1)
Usage:
  operator-sdk build <image> [flags]

Flags:
      --go-build-args string      Extra Go build arguments as one string such as "-ldflags -X=main.xyz=abc"
  -h, --help                      help for build
      --image-build-args string   Extra image build arguments as one string such as "--build-arg https_proxy=$https_proxy"
      --image-builder string      Tool to build OCI images. One of: [docker, podman, buildah] (default "docker")

Global Flags:
      --verbose   Enable verbose logging
```