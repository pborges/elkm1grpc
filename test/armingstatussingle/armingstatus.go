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
	as, err := elk.ArmingStatus(context.Background(), &elkm1grpc.ArmingStatusArgs{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("ArmingStatusReport -----------------------------------------")
	log.Println("ArmingStatus:", as.Areas[0].ArmingStatus.String())
	log.Println("ArmUpState  :", as.Areas[0].ArmUpState.String())
	log.Println("AlarmState  :", as.Areas[0].AlarmState.String())
}
