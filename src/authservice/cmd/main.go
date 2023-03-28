package main

import (
	"flag"

	"authservice/cmd/server"
)

var cfg = flag.String("config", "config/config.yaml", "config file location")

// main main
func main() {
	server.Run(*cfg)
}
