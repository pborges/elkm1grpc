package main

import (
	"os"
	"log"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"github.com/pborges/elkm1grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	c, err := net.Dial("tcp", "stratus:3000")
	if err != nil {
		panic(err)
	}

	elkServer := elkm1grpc.NewServer(c)

	elkm1grpc.RegisterElkGRPCServer(s, elkServer)

	log.Println("listening")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("going down")
}
