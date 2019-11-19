# Operator Sample for Che

## Instructions to run the demo

### Start the workspace

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

# Start the operator-hello-world workspace
DEVFILE=/Users/mloriedo/github/operator-hello-world/devfile.yaml
chectl workspace:start --devfile=${DEVFILE}
```

### Set Che user Preferences

```json
{
   "workbench.appearance.colorTheme": "light",
   "terminal.integrated.copyOnSelection": true,
   "git.user.name": "l0rd",
   "git.user.email": "mario.loriedo@gmail.com"
}
```

### Locally deploy and test the podset-operator

```bash
# Deploy the service account and setup RBAC
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml

# Setup the CRD and deploy the podset-operator
kubectl create -f deploy/crds/app.example.com_podsets_crd.yaml
kubectl create -f deploy/operator.yaml

# Create a podSet
echo "apiVersion: app.example.com/v1alpha1
kind: PodSet
metadata:
  name: example-podset
spec:
  replicas: 3" | oc create -f -
```

### Build the operator

```bash
# Using the operator SDK
operator-sdk build mariolet/podset-operator

# Go Build
go build -o /projects/src/github.com/l0rd/operator-hello-world/build/_output/bin/operator-hello-world \
         -gcflags all=-trimpath=/projects/src/github.com/l0rd \
         -asmflags all=-trimpath=/projects/src/github.com/l0rd \
         github.com/l0rd/operator-hello-world/cmd/manager
```

## Open Issues

### Blockers

- [ ] The operator, when deployed do not work as expected
- [ ] The build of the operator image does not work from Che

### Nice to have

- [ ] fix issue described in issue1.md
- [ ] run on CRW
- [ ] We should write down instructions to set imagePullStrategy to never for all workspaces components
- [ ] We should have a script to pre pull all needed images
- [ ] Setup the powershell PS1
- [ ] Setup the factory link
- [ ] upgrade operator-sdk to latest
