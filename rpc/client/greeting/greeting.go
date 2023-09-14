package greeting

import (
	v1 "go-gin-rpc/gen/greeting/v1"
	"go-gin-rpc/rpc/client"
)

func NewClient() (v1.GreetingClient, error) {
	conn, err := client.NewConn()

	return v1.NewGreetingClient(conn), err
}
