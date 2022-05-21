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
```