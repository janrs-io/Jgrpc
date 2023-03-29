package service

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/janrs-io/Jgrpc/src/pongservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1"
)

// NewClient New service's client
func NewClient(conf *config.Config) (v1.PongServiceClient, error) {

	serverAddress := conf.Grpc.Host + conf.Grpc.Port
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := v1.NewPongServiceClient(conn)
	return client, nil

}
