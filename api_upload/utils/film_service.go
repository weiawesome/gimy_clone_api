package utils

import (
	"google.golang.org/grpc"
)

var filmServiceConnection *grpc.ClientConn

func InitFilmServiceConnection() error {
	var err error
	filmServiceConnection, err = grpc.Dial(EnvFilmServiceAddress(), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	return err
}
func GetFilmServiceConnection() *grpc.ClientConn {
	return filmServiceConnection
}
func CloseFilmServiceConnection() error {
	return filmServiceConnection.Close()
}
