# Useful Commands

### Delete PODs prefix name

Prefix: goserver
```shell
kubectl get pods -n default --no-headers=true | awk '/goserver/{print $1}'| xargs  kubectl delete -n default pod
```

### Change Context/Account

```shell
kubectl config get-contexts

kubectl config use-context kind-fullcycle-k8s
```

### Access PODs

```shell
#List pods name
kubectl get pods    

NAME                        READY   STATUS    RESTARTS   AGE
goserver-7cd785cf9d-m7n2k   1/1     Running   0          2m21s
goserver-7cd785cf9d-tk7sf   1/1     Running   0          14h
mysql-0                     1/1     Running   0          13h
mysql-1                     1/1     Running   0          13h
mysql-2                     1/1     Running   0          13h

# Access by name (goserver-7cd785cf9d-tk7sf )
kubectl exec -it goserver-7cd785cf9d-tk7sf -- bash
```
