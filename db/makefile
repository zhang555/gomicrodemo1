build:
	docker build -t db .

run:
	docker run \
		-d \
		--name=db1 \
		--rm \
		--ip=10.0.0.200 \
		--net="net2" \
		-p 3336:3306 \
		-e MYSQL_ROOT_PASSWORD=88886666 \
		db

