package main

import (
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	port string
}

func NewgRPConfig(port string) *gRPCServer {
	return &gRPCServer{
		port: port,
	}
}

func (g *gRPCServer) NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func (g *gRPCServer) StartGRPCServer(s *grpc.Server) error {
	lis, err := net.Listen("tcp", g.port)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}