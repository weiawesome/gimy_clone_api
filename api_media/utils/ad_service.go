package utils

import (
	"google.golang.org/grpc"
)

var connection *grpc.ClientConn

func InitConnection() error {
	var err error
	connection, err = grpc.Dial(EnvAdServiceAddress(), grpc.WithInsecure(), grpc.WithBlock())
	return err
}
func GetConnection() *grpc.ClientConn {
	return connection
}
func CloseConnection() error {
	return connection.Close()
}
