# Nginx grpc servers lb example

```bash
# 1. start nginx + grpc servers
docker-compose up

# 2. build grpc client docker image
make client-build

# 3. send client message to grpc servers
make client-send  
make client-send
make client-send
make client-send
...
```

## Links
* https://www.youtube.com/watch?v=iPxhJWZ0IyU
* https://www.nginx.com/blog/deploying-nginx-plus-as-an-api-gateway-part-3-publishing-grpc-services/