# [FairwindsOps / goldilocks](https://github.com/FairwindsOps/goldilocks)

kubectl create namespace goldilocks
kubectl -n goldilocks apply -f kubernetes/autoscaler/vertical-pod-autoscaler/goldilocks-manifests/controller/
kubectl -n goldilocks apply -f kubernetes/autoscaler/vertical-pod-autoscaler/goldilocks-manifests/dashboard/

# Enable namespaces

kubectl label ns goldilocks goldilocks.fairwinds.com/enabled=true

# Viewing dashboard

kubectl -n goldilocks port-forward svc/goldilocks-dashboard 8080:80