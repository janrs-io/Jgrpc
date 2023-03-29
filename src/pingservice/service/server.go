package service

import (
	"context"
	"google.golang.org/grpc/grpclog"

	"github.com/janrs-io/Jgrpc/src/pingservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pingservice/genproto/v1"
	pongclientv1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1"
)

// Server Server struct
type Server struct {
	v1.UnimplementedPingServiceServer
	pingClient v1.PingServiceClient
	pongClient pongclientv1.PongServiceClient
	conf       *config.Config
}

// NewServer New service grpc server
func NewServer(
	conf *config.Config,
	pingClient v1.PingServiceClient,
	pongClient pongclientv1.PongServiceClient,
) v1.PingServiceServer {
	return &Server{
		pingClient: pingClient,
		pongClient: pongClient,
		conf:       conf,
	}
}

func (s *Server) Ping(ctx context.Context, req *v1.PingRequest) (*v1.PingResponse, error) {
	pongReq := &pongclientv1.PongRequest{Msg: "request from ping service"}
	pongResp, err := s.pongClient.Pong(ctx, pongReq)
	if err != nil {
		grpclog.Error("connect pong failed.[ERROR]=>" + err.Error())
		return nil, err
	}

	return &v1.PingResponse{
		Msg: "response ping msg:" + req.Msg + " and msg from pong service is: " + pongResp.Msg,
	}, nil
}
