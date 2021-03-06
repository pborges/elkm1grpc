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

	asClient, err := elk.ZoneChange(context.Background(), &elkm1grpc.ZoneChangeArgs{})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		z, err := asClient.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("ZoneChangeReport -----------------------------------------")
		log.Println("Zone:", z.Zone)
		log.Println("Description:", z.Description)
		log.Println("Status:", z.Status.String())
		log.Println("SubStatus:", z.SubStatus.String())
	}
}
