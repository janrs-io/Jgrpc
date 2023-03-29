//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"github.com/janrs-io/Jgrpc/src/pingservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pingservice/genproto/v1"
	"github.com/janrs-io/Jgrpc/src/pingservice/service"
)

// InitServer Inject service's component
func InitServer(conf *config.Config) (v1.PingServiceServer, error) {

	wire.Build(
		service.NewPongClient,
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil

}
