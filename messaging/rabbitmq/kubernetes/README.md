# RabbitMQ on kubernetes

## namespace: rabbit

kubectl apply -f messaging/rabbitmq/namespace.rabbitmq.yaml

## deploy the RabbitMQ as statefulset

kubectl apply -f messaging/rabbitmq/kubernetes

## access the management interface via portforwarding

kubectl -n rabbit port-forward rabbitmq-0 8080:15672
