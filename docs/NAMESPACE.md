[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
___
# Namespaces

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/38853ac660d4eb08f2a0d474617c13e51afb3d97)


```shell
# Create namespaces
kubectl create ns staging
kubectl create ns production

kubectl get ns              

NAME                 STATUS   AGE
default              Active   3d1h
kube-node-lease      Active   3d1h
kube-public          Active   3d1h
kube-system          Active   3d1h
local-path-storage   Active   3d1h
production           Active   10s
staging              Active   19s

# Create deployment by namespaces
kubectl apply -f k8s/namespaces/deployment.yaml -n=staging
kubectl apply -f k8s/namespaces/deployment.yaml -n=production

```

### Context

Create a context by namespace

```shell
kubectl delete deployments server-ns -n=staging
kubectl delete deployments server-ns -n=production

# View kubernetes configs 
kubectl config view

# Check current context
kubectl config current-context

# Create new context
kubectl config set-context kind-stg --namespace=staging --cluster=kind-fullcycle-k8s --user=kind-fullcycle-k8s
kubectl config set-context kind-prd --namespace=production --cluster=kind-fullcycle-k8s --user=kind-fullcycle-k8s


# Use context
kubectl config use-context kind-stg 
kubectl config use-context kind-prd

# Create deployment only in prd context
kubectl apply -f k8s/namespaces/deployment.yaml 
kubectl get deployments.apps
kubectl get po
```

___
[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
