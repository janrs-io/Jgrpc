//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"github.com/janrs-io/Jgrpc/src/pongservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1"
	"github.com/janrs-io/Jgrpc/src/pongservice/service"
)

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.PongServiceServer, error) {

	wire.Build(
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil

}
