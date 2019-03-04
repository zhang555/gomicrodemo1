build:
	protoc -I . --go_out=plugins=micro:$(GOPATH)/src/exercise/golibdemo/gomicrodemo/demo6/user-service proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service .

run:
#	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns  user-service
#        -e MICRO_BROKER_ADDRESS=0.0.0.0:4222 \
#        -e MICRO_BROKER=nats \

	docker run --net="host" \
        -p 50055 \
        -e MICRO_REGISTRY=mdns  \
	    -e DB_HOST= \
        -e DB_PASSWORD= \
        -e DB_NAME=micro \
        -e MICRO_BROKER_ADDRESS=0.0.0.0:4222 \
        -e MICRO_BROKER=nats \
        user-service

