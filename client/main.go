package main

import (
	"context"
	"fmt"
	"log"
	"afc-mano/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := proto.NewResourcesClient(conn)
	message := &proto.GetResourcesMessage{TargetResource: "cpu"}
	res, err := client.GetResources(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
