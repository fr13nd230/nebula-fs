package main

import (
	"net"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/alts"
)

type gRPCServer struct {
	port string
}

func NewConfig(port string) *gRPCServer {
	return &gRPCServer{
		port: port,
	}
}

func (g *gRPCServer) NewGrpcServer() *grpc.Server {
 //    altsCreds := alts.NewServerCreds(alts.DefaultServerOptions())
	// return grpc.NewServer(grpc.Creds(altsCreds))
	return grpc.NewServer()
}

func (g *gRPCServer) Start(grpcServer *grpc.Server) error {
	lis, err := net.Listen("tcp", g.port)
	if err != nil {
		return err
	}
	return grpcServer.Serve(lis)
}