#!/bin/sh

K8S_DIR="k8s/"

kubectl apply -f "${K8S_DIR}service.yaml"
kubectl apply -f "${K8S_DIR}configmap-env.yaml"
kubectl apply -f "${K8S_DIR}configmap-heroes.yaml"
kubectl apply -f "${K8S_DIR}metrics-server-non-tls.yaml"
kubectl apply -f "${K8S_DIR}secret.yaml"
kubectl apply -f "${K8S_DIR}persistent-volume.yaml"
kubectl apply -f "${K8S_DIR}hpa.yaml"
kubectl apply -f "${K8S_DIR}deployment.yaml"
