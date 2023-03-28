//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"authservice/config"
	v1 "authservice/genproto/v1"
	"authservice/service"
)

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.AuthServiceServer, error) {

	wire.Build(
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil

}
