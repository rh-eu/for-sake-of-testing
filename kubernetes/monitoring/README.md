# Kubernetes cluster monitoring with prometheus and grafana

## deploy prometheus-operator with kustomize

kubectl apply -f kubernetes/monitoring/namespace.monitoring.yaml

kubectl apply -k kubernetes/monitoring/kustomize/dev/
