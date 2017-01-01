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

	as, err := elk.Disarm(context.Background(), &elkm1grpc.ArmArgs{Area:1, Pin:1337})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("ArmingStatusReport -----------------------------------------")
	log.Println("ArmingStatus:", as.ArmingStatus.String())
	log.Println("ArmUpState  :", as.ArmUpState.String())
	log.Println("AlarmState  :", as.AlarmState.String())
}
