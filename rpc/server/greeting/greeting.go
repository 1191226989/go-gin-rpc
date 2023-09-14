package greeting

import (
	"context"
	v1 "go-gin-rpc/gen/greeting/v1"
)

func NewServer() *server {
	return &server{}
}

// server is used to implement greeting.GreetersServer
type server struct {
	v1.UnimplementedGreetingServer
}

// sayHello implements greeting.GreeterServer
func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}
