# Kubernetes cluster monitoring with Prometheus and Grafana

## Prometheus

kubectl apply -f kubernetes/monitoring/namespace.monitoring.yaml

### prometheus operator

kubectl apply -k kubernetes/monitoring/prometheus-operator/dev/

### cluster monitoring

kubectl apply -k kubernetes/monitoring/prometheus-cluster-monitoring/dev/

kubectl port-forward -n monitoring svc/prometheus-service 9090:9090

### metrics server

kubectl apply -k kubernetes/monitoring/metrics-server/dev/

### node exporter

kubectl apply -k kubernetes/monitoring/node-exporter/dev/

### kube state metrics

kubectl apply -k kubernetes/monitoring/kube-state-metrics/dev/

## Grafana

kubectl apply -k kubernetes/monitoring/grafana/dev/

kubectl port-forward -n monitoring svc/grafana 3000:3000
