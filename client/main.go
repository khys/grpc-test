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
	client := proto.NewResourceClient(conn)
	message := &proto.GetCpuMessage{Id: "0"}
	res, err := client.GetCpu(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
