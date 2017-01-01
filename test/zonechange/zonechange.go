package main

import (
	"os"
	"log"
	"github.com/pborges/elkm1grpc"
	"google.golang.org/grpc"
	"context"
)

func main() {
	log.SetOutput(os.Stdout)

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	elk := elkm1grpc.NewElkGRPCClient(conn)

	asClient, err := elk.ArmingStatusChange(context.Background(), &elkm1grpc.ArmingChangeArgs{})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		asr, err := asClient.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("ZoneChangeReport -----------------------------------------")
	}
}
