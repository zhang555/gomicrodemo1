build:
	docker build -t user-service .
run:
	docker run \
		--rm \
		--name=user-service1 \
		--net="net2" \
		-e MICRO_REGISTRY=mdns \
		-e DB_HOST=10.0.0.200 \
		-e DB_PASSWORD=88886666 \
		-e DB_NAME=micro  \
		user-service
