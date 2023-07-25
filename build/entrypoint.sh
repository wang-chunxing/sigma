#!/bin/sh

if [ "$(yq ".redis.type" < /etc/sigma/config.yaml)" = "internal"  ]; then
  if [ ! -d /var/lib/sigma/redis/ ]; then
    mkdir -p /var/lib/sigma/redis/
  fi
  redis-server /etc/sigma/redis.conf
  until nc -zv 127.0.0.1 6379; do echo waiting for redis; sleep 2; done
fi

exec "$@"