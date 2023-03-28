package service

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authclientv1 "authservice/genproto/v1"
	"permissionservice/config"
	v1 "permissionservice/genproto/v1"
)

func NewClient(conf *config.Config) (v1.PermissionServiceClient, error) {

	serverAddress := conf.Grpc.Host + conf.Grpc.Port
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := v1.NewPermissionServiceClient(conn)
	return client, nil

}

// NewAuthClient New auth service client
func NewAuthClient(conf *config.Config) (authclientv1.AuthServiceClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, conf.Client.Auth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("dial auth server failed.[ERROR]=>" + err.Error())
		return nil, err
	}
	client := authclientv1.NewAuthServiceClient(conn)
	return client, nil

}
