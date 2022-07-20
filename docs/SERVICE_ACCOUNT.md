[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
___
# Service Account

[Commit](https://github.com/silasstoffel/fullcyle-kubernetes/commit/43da5423b98c79aa12b8582ec574dd6b60bc40e1)


```shell
kubectl get serviceaccounts

kubectl apply -f k8s/namespaces/security.yaml

kubectl apply -f k8s/namespaces/deployment.yaml

kubectl get pods                                
NAME                         READY   STATUS    RESTARTS   AGE
server-ns-6f8cb9bd5c-gwr82   1/1     Running   0          2m57s

kubectl describe pod server-ns-6f8cb9bd5c-gwr82
```

___
[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
