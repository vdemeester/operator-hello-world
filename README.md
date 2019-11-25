# Operator Sample for Che

## Developer Workspace

[![Contribute](https://che.openshift.io/factory/resources/factory-contribute.svg)](https://che.openshift.io/f?url=https://raw.githubusercontent.com/l0rd/operator-hello-world/master/devfile.yaml)

This Che Factory can also be invoked with any host:
{hostURL}/f?url=https://github.com/l0rd/operator-hello-world/
It will read the `devfile.yaml` from the repository to instanciate the developer workspace.

## Instructions to run the demo

The following paragraphs describe how to setup and run a demo that shows how to develop this operator within Che.

The demo can be preceded by a short Che presentation ([slides](https://docs.google.com/presentation/d/19iAl2admu9fTrnVaTYxoDE43qvJ0ij9OM9NukcGyUDs/edit?usp=sharing))

### SETUP

#### Build the operator-sdk image

WARNING: the resulting image is > 4GB!! The build clones the operator source code and does a build of it. During the build all the go dependencies are downloaded. That saves a lot of time during the demo but makes the image uge.

WARNING: this should be done the day before the demo or the same day so that the image has all the operator dependencies already downloaded.

```bash
docker build -t mariolet/operator-sdk:0.11.0 -f build/Dockerfile.dev ./build
docker push mariolet/operator-sdk:0.11.0
```

Once you have build and push it you should replace the image reference in the `devfile.yaml`

Something that may be useful is to pull the image locally:

```bash
eval $(minikube docker-env)
docker pull mariolet/operator-sdk:<SHA256 DIGEST>
```

#### Start the k8s cluster (minikube) and start Che

```bash
# Start minikube
mk

# Start che
chectl server:start --platform=minikube --installer=helm

# Edit the run as user
kubens che
k edit cm/che # <- CHE_INFRA_KUBERNETES_POD_SECURITY__CONTEXT_FS__GROUP = 0
              #    CHE_INFRA_KUBERNETES_POD_SECURITY__CONTEXT_RUN__AS__USER = 0
k scale --replicas=0 deployment/che
k scale --replicas=1 deployment/che
```

#### Start a local workspace

```bash
# Start the operator-hello-world workspace
DEVFILE=/Users/mloriedo/github/operator-hello-world/devfile.yaml
chectl workspace:start --devfile=${DEVFILE}
```

#### Inject the kube config in the workspace

```bash
chectl workspace:inject -k -n <worksapce namespace>
```

#### Set Che user Preferences (optional)

```json
{
   "workbench.appearance.colorTheme": "light",
   "terminal.integrated.copyOnSelection": true,
   "git.user.name": "l0rd",
   "git.user.email": "mario.loriedo@gmail.com"
}
```

#### Locally deploy the podset-operator

```bash
# Deploy the service account and setup RBAC
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml

# Setup the CRD and deploy the podset-operator
kubectl create -f deploy/crds/app.example.com_podsets_crd.yaml
kubectl create -f deploy/operator.yaml
```

### DEMO

#### Show the podset operator

```bash
# Look for the podset api
kubectl api-resources

# Get the podsets (there should be no podset)
kubectl get podset

# Create a podSet
echo "apiVersion: app.example.com/v1alpha1
kind: PodSet
metadata:
  name: example-podset
spec:
  replicas: 3" | kubectl create -f -

# Get the podset again (there should be one)
kubectl get podset

# Get the pods (there should be a lot)
kubectl get pod

# Delete the podset
kubectl delete podset example-podset

# Scaled down the operator deployment
kubectl scale --replicas=0 deploy podset-operator
```

### Fix the bug using a Che wokrspace

#### Open the github repo of the operator, show the devfile and start the workspace using the factory

https://github.com/l0rd/operator-hello-world

#### Switch to the local workspace because che.openshift.io takes time to load and doesn't have enough resources

- Show the source code
- Build and run the operator using the corresponding task
- Change the code
- Rebuild and run and verify that the changes to the controller have been taken into account

### Old stuff

#### Build the operator image

```bash
# Using the operator SDK
operator-sdk build --image-builder buildah mariolet/podset-operator

# Using buildah
buildah bud --format=docker --isolation=chroot -f build/Dockerfile -t mariolet/podset-operator .

buildah bud --format=docker -f build/Dockerfile -t mariolet/podset-operator .

buildah bud --security-opt label=disable --security-opt seccomp=unconfined --format=docker --isolation=chroot -f build/Dockerfile -t mariolet/podset-operator .
```

#### Go build of the operator

```bash
go build -o /projects/src/github.com/l0rd/operator-hello-world/build/_output/bin/operator-hello-world \
         -gcflags all=-trimpath=/projects/src/github.com/l0rd \
         -asmflags all=-trimpath=/projects/src/github.com/l0rd \
         github.com/l0rd/operator-hello-world/cmd/manager
```

## Open Issues

### Blockers

- [ ] Investigate why the build of the operator image with buildah doesn't work in a Che container

### Nice to have

- [ ] fix issue described in issue1.md
- [ ] run on CRW
- [ ] We should write down instructions to set imagePullStrategy to never for all workspaces components
- [ ] We should have a script to pre pull all needed images
- [ ] Setup the powershell PS1
- [ ] upgrade operator-sdk to latest
