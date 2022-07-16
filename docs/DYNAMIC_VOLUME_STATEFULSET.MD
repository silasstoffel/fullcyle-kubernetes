[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
___
# ADDING DYNAMIC VOLUME WITH STATEFULSET

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/28cb8002d2c985ff53cb4245e3239d7aee0b44ed)


```shell
kubectl delete statefulsets mysql

kubectl apply -f k8s/statefulset.yaml

kubectl get pods

NAME                        READY   STATUS    RESTARTS   AGE
goserver-7cd785cf9d-99qlt   1/1     Running   0          31s
goserver-7cd785cf9d-ngpzw   1/1     Running   0          31s
mysql-0                     1/1     Running   0          10s
mysql-1                     1/1     Running   0          8s
mysql-2                     1/1     Running   0          5s
mysql-3                     1/1     Running   0          3s

kubectl get pvc

```

___
[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
