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

## Change Context

```shell
kubectl config use-context k8s-dev

kubectl config use-context kind-fullcycle-k8s
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

## Kubernetes API

To access API from Kubernetes you can run this command.

```shell
kubectl proxy --port=$port-number
# kubectl proxy --port=4001
```
Access: `http://localhost:${port-number}`

## Load Balancer

Create service type [load balancer](./k8s/service-load-balancer.yaml)

Apply

```shell
kubectl apply -f k8s/service-load-balancer.yaml 
```

In local environment it needs mapping port

```shell
kubectl port-forward service/goserver-load-balancer 8009:80
```

## Environment Vars

### Example 1
Commits:

[Definition env vars in service](https://github.com/silasstoffel/fullcyle-kubernetes/commit/d3a49ab5fb0e08f50253b2e2286f7d7489bfaee7)
[Load from env vars](https://github.com/silasstoffel/fullcyle-kubernetes/commit/f42177156a26ed5842e1ab217c68f7e5dd56c742)


```shell
# build new docker image
docker build -t silasstofel/hello-go:v2.0 -f ./docker/go/Dockerfile .

# push to docker hub
docker push silasstofel/hello-go:v2.0

# apply new deployment
kubectl apply -f k8s/deployment.yaml   
```

### Example 2 (config map)

Commits:

[Create config map](https://github.com/silasstoffel/fullcyle-kubernetes/commit/c0d2237241d28b83094c12b4a438fe1c3a9f46e6)

```shell
# apply config map
kubectl apply -f k8s/configmap-env.yaml

# apply new deployment
kubectl apply -f k8s/deployment.yaml   
```
### Example 3 (config map)

Commits:

[Change config map](https://github.com/silasstoffel/fullcyle-kubernetes/commit/d608f8a23be5d17d5a8e5977cb3f11d441ebce26)

```shell
# apply new deployment
kubectl apply -f k8s/deployment.yaml   
```

### Example 4 (config map with volumes)

[Commit Go source](https://github.com/silasstoffel/fullcyle-kubernetes/commit/5dad6b740246f0b5a91b7c48859ab8cf0a323c6c)

[Commit Create ConfigMap](https://github.com/silasstoffel/fullcyle-kubernetes/commit/5dad6b740246f0b5a91b7c48859ab8cf0a323c6c)

```shell
docker build -t silasstofel/hello-go:v3.1 -f docker/go/Dockerfile .

docker push silasstofel/hello-go:v3.1

kubectl apply -f k8s/configmap-heroes.yaml

kubectl apply -f k8s/deployment.yaml

kubectl port-forward service/goserver-service 8008:80
```

### Example 4 (Secrets)

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/8e38ac97edbbd29c281b1ec4c8bcf0867ecd7305)

```shell
docker build -t silasstofel/hello-go:v3.5 -f docker/go/Dockerfile .

docker push silasstofel/hello-go:v3.5

kubectl apply -f k8s/secret.yaml

kubectl apply -f k8s/deployment.yaml

kubectl port-forward service/goserver-service 8008:80
```

# Health Check

[Health Check](https://github.com/silasstoffel/fullcyle-kubernetes/commit/38f4d1824b08354104b86af0e96a21180ca7195c)

[Deployment](https://github.com/silasstoffel/fullcyle-kubernetes/commit/8597534743d3d13db7e0276b1d3106407062b9d9)

```shell
docker build -t silasstofel/hello-go:v3.7 -f docker/go/Dockerfile .

docker push silasstofel/hello-go:v3.7

kubectl apply -f k8s/deployment.yaml

kubectl port-forward service/goserver-service 8008:80
```

Check this endpoint http://localhost:8008/healthz, you should consider that 0-30 seconds is Ok (http code 200) and after 30 seconds the response is http code 500. 

# LivenessProbe

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/88501e3a09be49690efb33b0b3565c5bd15f15e8)

```shell
kubectl apply -f k8s/deployment.yaml

watch -n1 kubectl get pods

kubectl port-forward service/goserver-service 8008:80
```

# ReadinessProbe

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/708853669f0e8a912ae3956dc9c3b6eaa6ad70f5)

```shell
docker build -t silasstofel/hello-go:v3.8 -f docker/go/Dockerfile .

docker push silasstofel/hello-go:v3.8

kubectl apply -f k8s/deployment.yaml

kubectl port-forward service/goserver-service 8008:80
```

# ReadinessProbe and LivenessProbe Together

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/07faecef101184aea961db34f906404a20f6f750)

```shell
kubectl apply -f k8s/deployment.yaml

watch -n1 kubectl get pods

kubectl port-forward service/goserver-service 8008:80
```

# StartupProbe

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/9ebf61b7b90827b6a9f9ed2efc8d407f6c438ab5)

```shell
kubectl apply -f k8s/deployment.yaml

watch -n1 kubectl get pods

kubectl port-forward service/goserver-service 8008:80
```

# Change LivenessProbe, StartupProbe and ReadinessProbe in Application

Changed conditions to continue the course without resources of LivenessProbe, StartupProbe and ReadinessProbe.

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/91a1fab660e39a0b3b16563b626273158eb50013)
 
```shell
docker build -t silasstofel/hello-go:v3.9 -f docker/go/Dockerfile .

docker push silasstofel/hello-go:v3.9

kubectl apply -f k8s/deployment.yaml
```

# Install Metric Server

[Docs](https://github.com/kubernetes-sigs/metrics-server)

```shell
cd k8s

wget https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/high-availability.yaml

mv high-availability.yaml metrics-server-non-tls.yaml

# you must add argument --kubelet-insecure-tls for run in kind without TLS

cd ..

kubectl apply -f k8s/metrics-server-non-tls.yaml

# run this command to check out if it is available in API and find by v1beta1.metrics.k8s.io 
kubectl get apiservices
```
[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/eb46a391817d87f7a16d894ea78ec5ad53fdaae1)

# Resources

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/896cc2c1d9520e6a8533c6ed604d8a739cb390c4)

```shell
kubectl apply -f k8s/deployment.yaml

kubectl top pod goserver-654d9c9964-99fln
```

# HPA

[HPA file](https://github.com/silasstoffel/fullcyle-kubernetes/commit/cd938558982a8602f5c0381af0945a4f95ee7c8e)

```shell
kubectl apply -f k8s/hpa.yaml

kubectl get hpa
```

# Stress Test

[HPA file](https://github.com/silasstoffel/fullcyle-kubernetes/commit/b370a7671f72c206db6f9ceef4c75821efbe43da)

```shell
docker build -t silasstofel/hello-go:v3.10 -f docker/go/Dockerfile .

docker push silasstofel/hello-go:v3.10
```

Install [fortio](https://github.com/fortio/fortio)

Use must use host of the service (load balancer): http://goserver-service/healthz

```shell
# Run image with kubectl
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 900 -t 240s -c 80 "http://goserver-service/healthz"
```

Monitoring HPA

```shell
watch -n1 kubectl get pods

watch -n1 kubectl get hpa
```

# Persistent Volume

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/4ab4bf07c2ea7bf1c8def37d5689c22477ad6eac)

```shell
kubectl apply -f k8s/persistent-volume.yaml

kubectl get pvc
#or
kubectl get persistentvolumeclaims

#kubectl get storageclasses

kubectl apply -f k8s/deployment.yaml 

# To get volume name
kubectl get pvc

#NAME           STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
#goserver-pv1   Bound    pvc-3a87b670-3026-41e4-a280-8b2e69760252   2Gi        RWO            standard       23m

kubectl get pods

#NAME                        READY   STATUS    RESTARTS   AGE
#goserver-7cd785cf9d-fd5kr   1/1     Running   0          2m57s
#goserver-7cd785cf9d-zr7gx   1/1     Running   0          2m57s

kubectl exec -it goserver-7cd785cf9d-zr7gx -- bash

cd pvc1

touch custom-file.txt

exit

kubectl delete pod goserver-7cd785cf9d-zr7gx

kubectl get pods

#NAME                        READY   STATUS    RESTARTS   AGE
#goserver-7cd785cf9d-b78f4   1/1     Running   0          14s

kubectl exec -it goserver-7cd785cf9d-b78f4 -- bash

cd pvc1 

# must be able to show "custom-file.txt"
```

### Menu

- [StateFul + Headless](./docs/STATEFULSET_HEADLESS.md)
- [Dynamic volumes with Statefulset](./docs/DYNAMIC_VOLUME_STATEFULSET.md)
- [Ingress](./docs/INGRESS.md)
- [Namespace and Context](./docs/NAMESPACE.md)
- [Service Account](./docs/SERVICE_ACCOUNT.md)
- [Cluster Role](./docs/SERVICE_ACCOUNT.md)
