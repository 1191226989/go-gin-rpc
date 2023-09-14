rm -rf gen
protoc --go_out=. --go-grpc_out=. proto/greeting/v1/*.proto