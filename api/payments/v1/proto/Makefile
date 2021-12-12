regen-go:
	protoc --proto_path=. --proto_path=/home/dane/Documents/coding/googleapis \
	--go_out=../../ \
	--go-grpc_out=../../ \
	--grpc-gateway_out=../ \
	--grpc-gateway_opt=logtostderr=true \
    --grpc-gateway_opt=paths=source_relative \
	payment.proto
