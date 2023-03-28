//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"permissionservice/config"
	v1 "permissionservice/genproto/v1"
	"permissionservice/service"
)

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.PermissionServiceServer, error) {

	wire.Build(
		service.NewAuthClient,
		service.NewServer,
	)

	return &service.Server{}, nil

}
