# Apache Kafka

## [Quickstart](https://kafka.apache.org/quickstart)

### Build Kafka image from openjdk

cd messaging/kafka

docker build . -t mifomm/kafka:2.6.0

docker push mifomm/kafka:2.6.0

### Run Kafka broker

docker run --rm -d --name kafka-broker mifomm/kafka:2.6.0 && docker logs -f kafka-broker

### Create and describe a topic

docker run --rm -it --link kafka-broker mifomm/kafka:2.6.0 bin/kafka-topics.sh --create --topic mifomm --bootstrap-server kafka-broker:9092

docker run --rm -it --link kafka-broker mifomm/kafka:2.6.0 bin/kafka-topics.sh --describe --topic mifomm --bootstrap-server kafka-broker:9092

### Kafka producer and consumer

#### Write to a topic

docker run --rm -it --link kafka-broker --name kafka-producer mifomm/kafka:2.6.0 bin/kafka-console-producer.sh --topic mifomm --bootstrap-server kafka-broker:9092

#### Read from a topic

docker run --rm -it --link kafka-broker --name kafka-consumer mifomm/kafka:2.6.0 bin/kafka-console-consumer.sh --topic mifomm --from-beginning --bootstrap-server kafka-broker:9092
