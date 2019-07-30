package main

import (
	"log"
	"afc-mano/proto"
	"afc-mano/service"
	"net"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	resourceService := &service.ResourceService{}
	proto.RegisterResourcesServer(server, resourceService)
	server.Serve(listenPort)
}
