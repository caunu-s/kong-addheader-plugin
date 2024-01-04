# kong-addheader-plugin

```
docker cp ./kong-addheader-plugin docker-env-kong-kong-1:/usr/local/bin/kong-addheader-plugin
docker cp ./kong.conf docker-env-kong-kong-1:/etc/kong/kong.conf

docker compose --profile gateway restart

curl -i -X POST http://localhost:8001/plugins \
 --data name=kong-addheader-plugin

curl --head -X GET http://localhost:8000/mock/api/timezone/Asia/Tokyo
```