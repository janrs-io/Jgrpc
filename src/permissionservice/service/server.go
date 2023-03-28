package service

import (
	authv1 "authservice/genproto/v1"
	"context"
	"google.golang.org/grpc/grpclog"
	"permissionservice/config"
	v1 "permissionservice/genproto/v1"
)

// Server Server struct
type Server struct {
	v1.UnimplementedPermissionServiceServer
	authClient authv1.AuthServiceClient
	conf       *config.Config
}

// NewServer NewServer
func NewServer(
	conf *config.Config,
	authClient authv1.AuthServiceClient,
) v1.PermissionServiceServer {
	return &Server{
		authClient: authClient,
		conf:       conf,
	}
}

func (s Server) GetAdmin(ctx context.Context, req *v1.GetAdminRequest) (*v1.GetAdminResponse, error) {

	authReq := &authv1.PingRequest{Msg: "request from permission server"}
	authResp, err := s.authClient.Ping(ctx, authReq)
	if err != nil {
		grpclog.Error("connect auth failed.[ERROR]=>" + err.Error())
		return nil, err
	}
	if err != nil {
		grpclog.Error("get client failed.[ERROR]=>" + err.Error())
		return nil, err
	}

	return &v1.GetAdminResponse{
		Username: "user id is:" + req.Id,
		AuthMsg:  authResp.Msg,
	}, nil

}
