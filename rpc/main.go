package main

import (
	v1 "go-gin-rpc/gen/greeting/v1"
	"go-gin-rpc/rpc/server/greeting"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatalf("grpc failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterGreetingServer(s, greeting.NewServer())

	logrus.Printf("server listening at %v", listen.Addr())
	err = s.Serve(listen)
	if err != nil {
		logrus.Fatalf("rpc failed to serve: %v", err)
	}
}
