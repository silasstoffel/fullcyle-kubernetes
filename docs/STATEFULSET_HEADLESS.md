[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
___

# Statefulset and headless

[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/3e3544556e388f6f677951956b232d63b45cb121)

Pay attention: For it works, is necessary that `spec > serviceName: mysql-h` (statefulset) to be equal `metadata> name: mysql-h` (headless).

```shell
kubectl apply -f k8s/statefulset.yaml

kubectl apply -f k8s/mysql-headless-service.yaml

kubectl get services

NAME               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
goserver-service   ClusterIP   10.96.145.22   <none>        80/TCP     28m
kubernetes         ClusterIP   10.96.0.1      <none>        443/TCP    59m
mysql-h            ClusterIP   None           <none>        3306/TCP   36s
```

 
Following this steps to call specific PODs (mysql-0, mysql-1, mysql-2)

```shell
kubectl get pods    

NAME                        READY   STATUS    RESTARTS   AGE
goserver-7cd785cf9d-m7n2k   1/1     Running   0          2m21s
goserver-7cd785cf9d-tk7sf   1/1     Running   0          14h
mysql-0                     1/1     Running   0          13h
mysql-1                     1/1     Running   0          13h
mysql-2                     1/1     Running   0          13h

# Access POD via CLI
kubectl exec -it goserver-7cd785cf9d-tk7sf -- bash

# Access individual PODs of MySQL 
# Command model: ping ${pod_name}.${service_name}

ping mysql-0.mysql-h #master
ping mysql-1.mysql-h #slave1
ping mysql-2.mysql-h #slave2

```
[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
___
