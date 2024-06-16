cd api_consumer

docker build . -t go-api-consumer

cd ..

cd ws_publisher

docker build . -t go-ws-publisher

cd ..

cd nginx_reverse_proxy

docker build . -t nginx-reverse-proxy

cd ..

docker-compose up