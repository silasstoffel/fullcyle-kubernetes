#!/bin/sh

CLUSTER_NAME="fullcycle-k8s"
KUBCTL_CONTEXT="kind-fullcycle-k8s"

echo "Deleting cluster: ${CLUSTER_NAME}"

kind delete clusters $CLUSTER_NAME

echo "Creating cluster: ${CLUSTER_NAME}"
kind create cluster --config k8s/kind.yaml --name=${CLUSTER_NAME}


echo "Setting Cluster Context: ${KUBCTL_CONTEXT}"
kubectl cluster-info --context ${KUBCTL_CONTEXT}

echo "Nodes"
kubectl get nodes

echo 'Finish'
