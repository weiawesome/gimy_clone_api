package utils

import (
	"google.golang.org/grpc"
)

var adServiceConnection *grpc.ClientConn

func InitAdServiceConnection() error {
	var err error
	adServiceConnection, err = grpc.Dial(EnvAdServiceAddress(), grpc.WithInsecure(), grpc.WithBlock())
	return err
}
func GetAdServiceConnection() *grpc.ClientConn {
	return adServiceConnection
}
func CloseAdServiceConnection() error {
	return adServiceConnection.Close()
}
