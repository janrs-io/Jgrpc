package server

import (
	"authservice/config"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"

	v1 "authservice/genproto/v1/gw"
)

// RunHttpServer Run http server
func RunHttpServer(conf *config.Config) {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := v1.RegisterAuthServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		conf.Grpc.Port,
		opts,
	); err != nil {
		panic("register service handler failed.[ERROR]=>" + err.Error())
	}

	httpServer := &http.Server{
		Addr:    conf.Http.Port,
		Handler: mux,
	}
	fmt.Println("Listening http server on port" + conf.Http.Port)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			fmt.Println("listen http server failed.[ERROR]=>" + err.Error())
		}
	}()

	conf.Http.Server = httpServer

}
