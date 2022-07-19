[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
___
# INGRESS NGINX

[Commit 1](https://github.com/silasstoffel/fullcyle-kubernetes/commit/3e3544556e388f6f677951956b232d63b45cb121)

Ingress is a way to have a single external IP with many services.

[How install ingress in your cloud provider](https://kubernetes.github.io/ingress-nginx/deploy/)

This following example is using [helm](https://kubernetes.github.io/ingress-nginx/deploy/#quick-start)

[Create ingress file](https://github.com/silasstoffel/fullcyle-kubernetes/commit/45c9ffd4228849aebec75099b8a198c3affa5cb7)

```shell
kubectl apply -f k8s/ingress.yaml

kubectl get service

```

___
[Go Back](https://github.com/silasstoffel/fullcyle-kubernetes)
