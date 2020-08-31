# Vertical Pod Autoscaler

[Read on github](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler)

## Prerequisites

The metrics server must be installed in the cluster (kubectl get APIservice | grep metrics)

## Installation

kubectl apply -f kubernetes/autoscaler/vertical-pod-autoscaler/deploy/
