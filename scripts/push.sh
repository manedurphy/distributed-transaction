#! /bin/bash

APIGATEWAY_VERSION="1.0.5"
CUSTOMERS_VERSION="1.0.3"
ORDERS_VERSION="1.0.3"
PAYMENTS_VERSION="1.0.3"

docker buildx build \
	--platform linux/amd64,linux/arm/v7 \
	--tag manedurphy/distributed-transaction-apigateway:$APIGATEWAY_VERSION \
	--push \
	--file Dockerfile.apigateway .

docker buildx build \
	--platform linux/amd64,linux/arm/v7 \
	--tag manedurphy/distributed-transaction-customers:$CUSTOMERS_VERSION \
	--push \
	--file Dockerfile.customers .

docker buildx build \
	--platform linux/amd64,linux/arm/v7 \
	--tag manedurphy/distributed-transaction-orders:$ORDERS_VERSION \
	--push \
	--file Dockerfile.orders .

docker buildx build \
	--platform linux/amd64,linux/arm/v7 \
	--tag manedurphy/distributed-transaction-payments:$PAYMENTS_VERSION \
	--push \
	--file Dockerfile.payments .
