# RabbitMQ Clustering

## nodes: rabbit-1@rabbit-1, rabbit-2@rabbit-2, rabbit-3@rabbit-3

## provide the Erlang-Cookie - it *must* be same for each rabbit-node in order to build a cluster

## management-version with UI enabled, expose port 15672

docker run -d --rm --net rabbits -p 8081:15672 -v ${HOME}/erlang/:/var/lib/rabbitmq/ --hostname rabbit-1 --name rabbit-1 rabbitmq:3.8-management && docker logs -f rabbit-1

docker run -d --rm --net rabbits -p 8082:15672 -v ${HOME}/erlang/:/var/lib/rabbitmq/ --hostname rabbit-2 --name rabbit-2 rabbitmq:3.8-management && docker logs -f rabbit-2

docker run -d --rm --net rabbits -p 8083:15672 -v ${HOME}/erlang/:/var/lib/rabbitmq/ --hostname rabbit-3 --name rabbit-3 rabbitmq:3.8-management && docker logs -f rabbit-3

### check the cluster status

docker exec -it rabbit-1 rabbitmqctl cluster_status

## join rabbit-2

docker exec -it rabbit-2 rabbitmqctl stop_app
docker exec -it rabbit-2 rabbitmqctl reset
docker exec -it rabbit-2 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec -it rabbit-2 rabbitmqctl start_app
docker exec -it rabbit-2 rabbitmqctl cluster_status

## join rabbit-3

docker exec -it rabbit-3 rabbitmqctl stop_app
docker exec -it rabbit-3 rabbitmqctl reset
docker exec -it rabbit-3 rabbitmqctl join_cluster rabbit@rabbit-1
docker exec -it rabbit-3 rabbitmqctl start_app
docker exec -it rabbit-3 rabbitmqctl cluster_status

## publishing a message

### golang publisher

cd messaging/rabbitmq/golang-apps/publischer

docker build . -t mifomm/rabbitmq-publisher:1.0.0
docker push mifomm/rabbitmq-publisher:1.0.0

docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 8000:80 mifomm/rabbitmq-publisher:1.0.0

#### publish a message

curl -X POST http://localhost:8000/publish/the-new-message

### golang consumer

cd messaging/rabbitmq/golang-apps/consumer

docker build . -t mifomm/rabbitmq-consumer:1.0.0
docker push mifomm/rabbitmq-consumer:1.0.0

docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest mifomm/rabbitmq-consumer:1.0.0
