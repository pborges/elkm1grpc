package main

import (
	"os"
	"log"
	"net"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"github.com/pborges/elkm1grpc"
	"github.com/pkg/term"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	if len(os.Args) < 3 {
		log.Printf("usage: %s <serial port> <baud>\n", os.Args[0])
		os.Exit(1)
	}

	serialPortName := os.Args[1]
	baud, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println("[ERROR] Error parsing baud")
		os.Exit(1)
	}

	serialPort, err := term.Open(serialPortName)
	if err != nil {
		log.Printf("[ERROR] Unable to open Serial Port %+v\n", err)
		os.Exit(1)
	}
	serialPort.SetRaw()
	serialPort.SetSpeed(baud)

	s := grpc.NewServer()
	reflection.Register(s)

	elkServer := elkm1grpc.NewServer(serialPort)

	elkm1grpc.RegisterElkGRPCServer(s, elkServer)

	log.Println("listening")
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("going down")
}
