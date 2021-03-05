package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/kalmecak/grpcando/chat"
)

func main() {

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {

		log.Fatalln("Hubo un error prro: ", err.Error())
		return
	}

	grpcServer := grpc.NewServer()

	s := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {

		log.Fatalln("No sirvió tu mamada de gRPC: ", err.Error())
	}

	log.Println("Escuchando ladridos de los prros en el puerto :9000")
}
