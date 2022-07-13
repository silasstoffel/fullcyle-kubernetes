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
