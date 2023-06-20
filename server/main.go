package main

import (
	"fmt"
	"github.com/RakhimovAns/GRPC/pkg"
	pb "github.com/RakhimovAns/GRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	pkg.DataBaseConnection()
}

func main() {
	fmt.Println("grpc server running ...")
	lis, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatal("Failed to listen")
	}
	s := grpc.NewServer()
	pb.RegisterMovieServiceServer(s, &pkg.Server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
