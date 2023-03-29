package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/janrs-io/Jgrpc/src/pingservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pingservice/genproto/v1"
)

// RunGrpcServer Run grpc server
func RunGrpcServer(server v1.PingServiceServer, conf *config.Config) {

	grpcServer := grpc.NewServer()
	v1.RegisterPingServiceServer(grpcServer, server)

	fmt.Println("Listening grpc server on port" + conf.Grpc.Port)
	listen, err := net.Listen("tcp", conf.Grpc.Port)
	if err != nil {
		panic("listen grpc tcp failed.[ERROR]=>" + err.Error())
	}

	go func() {
		if err = grpcServer.Serve(listen); err != nil {
			log.Fatal("grpc serve failed", err)
		}
	}()

	conf.Grpc.Server = grpcServer

}
