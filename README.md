# FullCycle Kubernetes

## Install

- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Kind](https://kind.sigs.k8s.io/)
- [Docker](https://docs.docker.com/engine/install/)

## Create Cluster

```shell
kind create cluster

#connect on cluster
kubectl cluster-info --context kind-kind

#check nodes
kubectl get nodes 
```

## Delete Cluster

```shell
#list available clusters
kind get clusters

# delete
kind delete clusters $cluster-name
```

## Create Custom Node

```shell
kind create cluster --config k8s/kind.yaml --name=fullcycle-k8s

#connect cluster
kubectl cluster-info --context kind-fullcycle-k8s

#list nodes
kubectl get nodes
```

## Create image (golang) and push dockerhub

```shell
docker build -t silasstofel/hello-go -f docker/go/Dockerfile .

docker run --rm -p 8080:80 silasstofel/hello-go

docker push silasstofel/hello-go

#docker push silasstofel/hello-go:v1.0
#docker push silasstofel/hello-go:v1.1
#docker push silasstofel/hello-go:latest
```

## Create POD
```shell
kubectl apply -f k8s/pod.yaml

#check PODs
kubectl get pods  
```

## ReplicaSet
```shell
kubectl apply -f k8s/replicaset.yaml 

#check PODs
kubectl get pods

#list Replica Sets
kubectl get replicasets

#delete Replica Set
kubectl delete replicaset $name
#kubectl delete replicaset goserver
```

## Deployment

```shell
# Create/Apply
kubectl apply -f k8s/deployment.yaml 

#Check PODs
kubectl get pods

#Check Deployments
kubectl get deployments
```

Change version of docker image in ./k8s/deployment.yaml and apply new deployment with this command:

```shell
kubectl apply -f k8s/deployment.yaml
```

It will be create new replica set and pods. The old replica set and PODs will be replace.

```shell
kubectl apply -f k8s/deployment.yaml
```

For you check the used image, use this command to list information from POD.


```shell
#kubectl describe pod ${pod-name}
kubectl describe pod goserver-79457477d4-xr422
```
Now, find by `image` attribute to check image version used.

## Rollout Deployment

To get versions/revision your deployments

```shell
#kubectl rollout history deployment $deployment-name

kubectl rollout history deployment goserver
```

Use this command to Rollout / Rollback deployment

```shell
#kubectl rollout undo deployment $deployment-name [--to-revision=${revision-number}]

kubectl rollout undo deployment goserver

#check image version in current version
kubectl describe pod ${pod-name}

# list replica set
kubectl get replicaset
```

## Service: Cluster IP

See the service definition [here](./k8s/service.yaml)

```shell
kubectl apply -f k8s/service.yaml 
```

How view my started services? Run this command.

```shell
kubectl get services
# or
kubectl get svc
```

How access my application? We should use port redirect

```shell
# service/goserver-service = name of service
kubectl port-forward service/goserver-service 8008:80
#kubectl port-forward svc/goserver-service 8008:80
```

## Service: With param targetPort

1. In [server.go](./server.go) file, change port of your application is running
2. In [service.yaml](./k8s/service.yaml) file, add `targetPort` params with port of your container/application

Run this command:

```shell
kubectl port-forward service/goserver-service 8008:80
```
To check, you can access this address http://localhost:8008
