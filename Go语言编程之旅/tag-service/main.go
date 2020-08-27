package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/wsloong/tag-service/proto"
	"github.com/wsloong/tag-service/server"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":8009")
	if err != nil {
		log.Fatalf("net.Listen err: %s", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server.Serve err: %s", err)
	}
}