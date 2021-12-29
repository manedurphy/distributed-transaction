#! /bin/bash

docker buildx build \
	--platform linux/amd64 \
	--tag manedurphy/distributed-transaction-apigateway \
	--load \
	--file Dockerfile.apigateway .

docker buildx build \
	--platform linux/amd64 \
	--tag manedurphy/distributed-transaction-customers \
	--load \
	--file Dockerfile.customers .

docker buildx build \
	--platform linux/amd64 \
	--tag manedurphy/distributed-transaction-orders \
	--load \
	--file Dockerfile.orders .

docker buildx build \
	--platform linux/amd64 \
	--tag manedurphy/distributed-transaction-payments \
	--load \
	--file Dockerfile.payments .
