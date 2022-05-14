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

# Create Custom Node

```shell
kind create cluster --config k8s/kind.yaml --name=fullcycle-k8s

#connect cluster
kubectl cluster-info --context kind-fullcycle-k8s

#list nodes
kubectl get nodes
```

# Create image (golang) and push dockerhub

```shell
docker build -t silasstofel/hello-go -f docker/go/Dockerfile .

docker run --rm -p 8080:80 silasstofel/hello-go

docker push silasstofel/hello-go
```

# Create POD
```shell
kubectl apply -f k8s/pod.yaml

#check PODs
kubectl get pods  
```

# ReplicaSet
```shell
kubectl apply -f k8s/replicaset.yaml 

#check PODs
kubectl get pods

#list replicasets
kubectl get replicasets

```