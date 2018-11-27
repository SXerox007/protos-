package main

import (
	"log"
	"net"
	"protos-/expert/services/proto"

	"google.golang.org/grpc"
)

func main() {
	Init()
}

func Init() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error in server listen: ", listner)
	}

	srv := grpc.NewServer()
	//To register the server in gRPC
	search.RegisterCreateServiceServer(srv, &Server{})

	if err := srv.Serve(listner); err != nil {
		log.Fatal("Error in Serve the Server:", err)
	}
}

type Server struct {
}
