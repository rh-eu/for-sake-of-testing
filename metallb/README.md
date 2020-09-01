# [MetalLB](https://github.com/metallb/metallb)

## applying v0.9.3 to Kind-cluster

kubectl apply -f metallb/0.9.3/manifests/namespace.yaml

kubectl apply -f metallb/0.9.3/manifests/metallb.yaml

### on first install only

kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)"

### configmap

kubectl apply -f metallb/0.9.3/manifests/config.yaml
