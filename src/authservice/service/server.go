package service

import (
	"authservice/config"
	"context"

	v1 "authservice/genproto/v1"
)

// Server Server struct
type Server struct {
	v1.UnimplementedAuthServiceServer
	authClient v1.AuthServiceClient
	conf       *config.Config
}

// NewServer NewServer
func NewServer(conf *config.Config, authClient v1.AuthServiceClient) v1.AuthServiceServer {
	return &Server{
		authClient: authClient,
		conf:       conf,
	}
}

func (s *Server) Ping(ctx context.Context, req *v1.PingRequest) (*v1.PongResponse, error) {
	return &v1.PongResponse{Msg: "response msg:" + req.Msg}, nil
}
