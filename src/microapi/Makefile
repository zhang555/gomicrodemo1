build:
	docker build -t demo2 .

run:
#	docker run \
#		microapi \
#		api --handler=rpc     --address=:8080   --namespace=gomicrodemo1  --cors=http://localhost

	docker run \
		-e MICRO_REGISTRY=mdns \
		microapi api
