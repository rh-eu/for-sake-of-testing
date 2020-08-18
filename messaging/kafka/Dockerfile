FROM openjdk
RUN cd /opt && curl -OL https://apache.mirror.digionline.de/kafka/2.6.0/kafka_2.13-2.6.0.tgz && tar -zxvf kafka_2.13-2.6.0.tgz && rm kafka_2.13-2.6.0.tgz
WORKDIR /opt/kafka_2.13-2.6.0/
COPY start.sh /opt/kafka_2.13-2.6.0/
CMD ["./start.sh"]
