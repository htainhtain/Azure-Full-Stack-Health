FROM redis:alpine

WORKDIR /data

RUN adduser -D myredis && chown -R myredis:myredis /data

USER redis

EXPOSE 6379