# Redis

## [Redis configuration](https://redis.io/topics/config)

### [Redis Docker Images](https://hub.docker.com/_/redis)

#### Starting redis with custom config

cd messages/redis

docker network create redis

docker run --rm -it --name redis --net redis -v ${PWD}/config/:/etc/redis/ redis:6.0-alpine redis-server /etc/redis/redis.conf
