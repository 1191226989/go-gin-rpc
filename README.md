# go-gin-rpc
go-gin-rpc for clean architecture

1. Install protoc
download zip package，unzip and copy `/bin/protoc` to the directory `/usr/local/bin`

https://github.com/protocolbuffers/protobuf/releases/download/v24.3/protoc-24.3-linux-x86_64.zip


2. Install the protocol compiler plugins for Go using the following commands:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

3. Update your PATH so that the protoc compiler can find the plugins:
```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

4. Regenerate gRPC code
```sh
protoc --go_out=gen --go_opt=paths=source_relative \
  --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
  -I=$PWD pb/greeting/v1/greeting.proto
```
or
```sh
./build.sh
```

5. Running

First Step: run grpc server
```sh
go run rpc/main.go
```
Second Step: run gin server
```sh
go run api/main.go
```

6. Testing
Send data to gin server
```sh
curl -v 'http://localhost:8080/api/greeting/gin'
```

#### 目录结构
```
.
├── api                         # gin http api
│   ├── configs
│   ├── internal
│   │   ├── controller
│   │   │   ├── article
│   │   │   │   └── detail.go
│   │   │   └── greeting
│   │   │       └── hello.go
│   │   └── router
│   │       └── router.go
│   └── main.go                 # api server
├── build.sh                    # 编译工具
├── gen                         # protoc 编译目录
│   └── greeting
│       └── v1
│           ├── greeting_grpc.pb.go
│           └── greeting.pb.go
├── go.mod
├── go.sum
├── proto                       # proto buffer 文件
│   └── greeting
│       └── v1
│           └── greeting.proto
├── README.md
└── rpc                         # rpc 服务端和客户端
    ├── client
    │   ├── client.go
    │   └── greeting
    │       └── greeting.go
    ├── configs
    ├── internal
    │   ├── config
    │   └── model
    ├── main.go                 # rpc 服务端 server
    └── server
        └── greeting
            └── greeting.go

```


