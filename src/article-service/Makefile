build:
	docker build -t article-service .

run:
    docker run --net="host" \
		-p 5054 \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=192.168.0.200 \
		-e DB_PASSWORD=$dbpassword \
		-e DB_NAME=micro  \
		-e MICRO_BROKER=nats  \
		-e MICRO_BROKER_ADDRESS=0.0.0.0:4222   \
		article-service

