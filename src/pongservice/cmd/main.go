package main

import (
	"flag"

	"github.com/janrs-io/Jgrpc/src/pongservice/cmd/server"
)

var cfg = flag.String("config", "config/config.yaml", "config file location")

// main main
func main() {
	server.Run(*cfg)
}
