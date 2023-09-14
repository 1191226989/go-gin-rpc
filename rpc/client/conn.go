package client

import (
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	err  error
	conn *grpc.ClientConn
	once sync.Once
)

// 单例
func NewConn() (*grpc.ClientConn, error) {
	once.Do(func() {
		conn, err = grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logrus.Fatalf("did not connect: %v", err)
		}
		// defer conn.Close()
	})
	return conn, err
}
