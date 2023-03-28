package main

import (
	"flag"
	"permissionservice/cmd/server"
)

var cfg = flag.String("config", "config/config.yaml", "database config file location")

// main main
func main() {
	server.Run(*cfg)
}
