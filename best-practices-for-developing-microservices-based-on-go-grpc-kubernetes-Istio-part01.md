[Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part01](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part01.md "Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part01")

[Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part02](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part02.md "Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part02")

[Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part03](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part03.md "Best practices for developing microservices based on Go/Grpc/kubernetes/Istio - Part03")

***

# [Jgrpc](https://github.com/janrs-io/Jgrpc "Jgrpc")

This [project](https://github.com/janrs-io/Jgrpc "project") provides a reference for best practices for developing
microservices based on `Go/Grpc/kubernetes/Istio`.

And it implements `CICD` based on `Jenkins/Gitlab/Harbor`.

And use the [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway "grpc-gateway") as a gateway proxy.

This best practice is divided into three parts:

1. [Create a microservice named `pongservice`](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part01.md "Create a microservice named `pongservice`")
2. [Create a microservice named `pingservice` and access `pongservice`](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part02.md "Create a microservice named `pingservice` and access `pongservice`")
3. [Create a `CICD` deployment process base on `Jenkins/Gitlab/Harbor`](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part03.md "Create a `CICD` deployment process base on `Jenkins/Gitlab/Harbor`")

In this part, we will create the `pongservice` microservice.

***

# Prerequisites

It is assumed that the following tools have been installed:

- `protoc-gen-grpc-gateway`
- `protoc-gen-openapiv2`
- `protoc-gen-go`
- `protoc-gen-go-grpc`
- `buf`
- `wire`

Here are the tutorial addresses for installing these tools:

`protoc-gen-grpc-gateway` & `protoc-gen-openapiv2` & `protoc-gen-go` & `protoc-gen-go-grpc`
For installation tutorials of these four tools, please
view:[install protoc-gen* tools](https://github.com/grpc-ecosystem/grpc-gateway#compile-from-source "install protoc-gen* tools")

`wire`
For installation tutorials of this tools, please
view:[install wire tool](https://github.com/google/wire#installing "install wire tool")

`buf`
For installation tutorials of this tools, please
view:[install buf tool](https://buf.build/docs/installation/ "install buf tool")

# Project Structure

The final structure of the directory in this part is as follows:

```bash
Jgrpc
├── devops
├── istio-manifests
├── kubernetes-manifests
└── src
    └── pongservice
        ├── buf.gen.yaml
        ├── cmd
        │   ├── main.go
        │   └── server
        │       ├── grpc.go
        │       ├── http.go
        │       ├── run.go
        │       ├── wire.go
        │       └── wire_gen.go
        ├── config
        │   ├── config.go
        │   └── config.yaml
        ├── genproto
        │   └── v1
        │       ├── gw
        │       │   └── pongservice.pb.gw.go
        │       ├── pongservice.pb.go
        │       └── pongservice_grpc.pb.go
        ├── go.mod
        ├── go.sum
        ├── proto
        │   ├── buf.lock
        │   ├── buf.yaml
        │   └── v1
        │       ├── pongservice.proto
        │       └── pongservice.yaml
        └── service
            ├── client.go
            └── server.go

14 directories, 20 files
```

***

# Start

Create the overall directory structure of the project as follows:

```bash
Jgrpc
├── devops
├── istio-manifests
├── kubernetes-manifests
├── src
```

# Create pongservice microservice

## Create base directory

Create a microservice named `pongservice` in the `src` directory.The directory structure is as follows:

```bash
pongservice
├── cmd
│   └── server
├── config
├── proto
│   └── v1
└── service

6 directories, 0 files
```

## Generate codes & files

**IN THIS STEP YOU NEED TO FOCUS ON.**

### Generate buf.yaml

Execute the following command in the `pongservice/proto` directory:

```bash
buf mod init
```

This command creates a file named `buf.yaml` located in the `ponservice/proto` directory.
The code of `buf.yaml` is as follows:

```yaml
version: v1
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
```

Add the following dependency code between `version` and `breaking`:

```yaml
deps:
  - buf.build/googleapis/googleapis
```

Please see the reason for adding this dependency
code:[https://github.com/grpc-ecosystem/grpc-gateway#3-external-configuration](https://github.com/grpc-ecosystem/grpc-gateway#3-external-configuration "Generate reverse-proxy using protoc-gen-grpc-gateway")

The complete `buf.yaml` file after adding the dependent code is as follows:

```yaml
version: v1
deps:
  - buf.build/googleapis/googleapis
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
```

Then execute the following command in the `pongservice/proto` directory:

```bash
buf mod update
```

After executing the command, a `buf.lock` file will be generated, the code is as follows:

```lock
# Generated by buf. DO NOT EDIT.
version: v1
deps:
  - remote: buf.build
    owner: googleapis
    repository: googleapis
    commit: 463926e7ee924d46ad0a726e1cf4eacd
```

### Generate pongservice.proto

Create a proto file named `pongservice.proto` with the following code in the `pongservice/proto/v1` directory:

```proto
syntax = "proto3";

package proto.v1;

option go_package = "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1";

service PongService {
  rpc Pong(PongRequest) returns(PongResponse){}
}

message PongRequest {
  string msg = 1 ;
}

message PongResponse {
  string msg = 1;
}
```

### Generate pongservice.yaml

Create a proto file named `pongservice.yaml` with the following code in the `pongservice/proto/v1` directory:

```yaml
type: google.api.Service
config_version: 3

http:
  rules:
    - selector: proto.v1.PongService.Pong
      get: /pong.v1.pong
```

### Generate buf.gen.yaml

Create a yaml file named `buf.gen.yaml` with the following code in the `pongservice` directory:

```yaml
version: v1
plugins:
  - plugin: go
    out: genproto/v1
    opt:
      - paths=source_relative
  - plugin: go-grpc
    out: genproto/v1
    opt:
      - paths=source_relative
  - plugin: grpc-gateway
    out: genproto/v1/gw
    opt:
      - paths=source_relative
      - grpc_api_configuration=proto/v1/pongservice.yaml
      - standalone=true
```

In the `pongservice` directory, execute the following command:

```bash
buf generate proto/v1
```

After executing the command, a genproto directory will be automatically created in the `pongservice` directory, and
there are the following files in this directory:

```bash
genproto
└── v1
    ├── gw
    │   └── ponservice.pb.gw.go
    ├── ponservice.pb.go
    └── ponservice_grpc.pb.go

2 directories, 3 files
```

### Generate go.mod

Create `go.mod` in the `pongservice` directory:

```bash
go mod init github.com/janrs-io/Jgrpc/src/pongservice && go mod tidy
```

### Generate config.yaml

In the `pongservice/config` directory, create `config.yaml` file and add the following code:

```yaml
# grpc config
grpc:
  host: ""
  port: ":50051"
  name: "pong-grpc"

# http config
http:
  host: ""
  port: ":9001"
  name: "pong-http"
```

### Generate config.go

In the `pongservice/config` directory, create `config.go` file and add the following code:

```go
package config

import (
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Config Service config
type Config struct {
	Grpc Grpc `json:"grpc" yaml:"grpc"`
	Http Http `json:"http" yaml:"http"`
}

// NewConfig Initial service's config
func NewConfig(cfg string) *Config {

	if cfg == "" {
		panic("load config file failed.config file can not be empty.")
	}

	viper.SetConfigFile(cfg)

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		panic("read config failed.[ERROR]=>" + err.Error())
	}
	conf := &Config{}
	// Assign the overloaded configuration to the global
	if err := viper.Unmarshal(conf); err != nil {
		panic("assign config failed.[ERROR]=>" + err.Error())
	}

	return conf

}

// Grpc Grpc server config
type Grpc struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Name   string `json:"name" yaml:"name"`
	Server *grpc.Server
}

// Http Http server config
type Http struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Name   string `json:"name" yaml:"name"`
	Server *http.Server
}
```

Then run `go mod tidy` again in the `pongservice` directory.

### Generate client.go

In the `pongservice/service` directory, create `client.go` file and add the following code:

```go
package service

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/janrs-io/Jgrpc/src/pongservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1"
)

// NewClient New service's client
func NewClient(conf *config.Config) (v1.PongServiceClient, error) {

	serverAddress := conf.Grpc.Host + conf.Grpc.Port
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := v1.NewPongServiceClient(conn)
	return client, nil

}
```

In the `pongservice/service` directory, create `server.go` file and add the following code:

```go
package service

import (
	"context"

	"github.com/janrs-io/Jgrpc/src/pongservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1"
)

// Server Server struct
type Server struct {
	v1.UnimplementedPongServiceServer
	pongClient v1.PongServiceClient
	conf       *config.Config
}

// NewServer New service grpc server
func NewServer(conf *config.Config, pongClient v1.PongServiceClient) v1.PongServiceServer {
	return &Server{
		pongClient: pongClient,
		conf:       conf,
	}
}

func (s *Server) Pong(ctx context.Context, req *v1.PongRequest) (*v1.PongResponse, error) {
	return &v1.PongResponse{Msg: "response pong msg:" + req.Msg}, nil
}
```

### Generate run server files

In the `pongservice/cmd/server` directory, create the following four files:

- `grpc.go`
- `http.go`
- `run.go`
- `wire.go`

Add the following code to the `grpc.go` file:

```go
package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/janrs-io/Jgrpc/src/pongservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1"
)

// RunGrpcServer Run grpc server
func RunGrpcServer(server v1.PongServiceServer, conf *config.Config) {

	grpcServer := grpc.NewServer()
	v1.RegisterPongServiceServer(grpcServer, server)

	fmt.Println("Listening grpc server on port" + conf.Grpc.Port)
	listen, err := net.Listen("tcp", conf.Grpc.Port)
	if err != nil {
		panic("listen grpc tcp failed.[ERROR]=>" + err.Error())
	}

	go func() {
		if err = grpcServer.Serve(listen); err != nil {
			log.Fatal("grpc serve failed", err)
		}
	}()

	conf.Grpc.Server = grpcServer

}
```

Add the following code to the `http.go` file:

```go
package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/janrs-io/Jgrpc/src/pongservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1/gw"
)

// RunHttpServer Run http server
func RunHttpServer(conf *config.Config) {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := v1.RegisterPongServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		conf.Grpc.Port,
		opts,
	); err != nil {
		panic("register service handler failed.[ERROR]=>" + err.Error())
	}

	httpServer := &http.Server{
		Addr:    conf.Http.Port,
		Handler: mux,
	}
	fmt.Println("Listening http server on port" + conf.Http.Port)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			fmt.Println("listen http server failed.[ERROR]=>" + err.Error())
		}
	}()

	conf.Http.Server = httpServer

}
```

Add the following code to the `run.go` file:

```go
package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/janrs-io/Jgrpc/src/pongservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1"
)

// Run Run service server
func Run(cfg string) {

	conf := config.NewConfig(cfg)
	// run grpc server
	RunGrpcServer(initServer(conf), conf)
	// run http server
	RunHttpServer(conf)
	// listen exit server event
	HandleExitServer(conf)

}

// SetServer Wire inject service's component
func initServer(conf *config.Config) v1.PongServiceServer {
	server, err := InitServer(conf)
	if err != nil {
		panic("run server failed.[ERROR]=>" + err.Error())
	}
	return server
}

// HandleExitServer Handle service exit event
func HandleExitServer(conf *config.Config) {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conf.Grpc.Server.GracefulStop()
	if err := conf.Http.Server.Shutdown(ctx); err != nil {
		panic("shutdown service failed.[ERROR]=>" + err.Error())
	}
	<-ctx.Done()
	close(ch)
	fmt.Println("Graceful shutdown http & grpc server.")

}
```

Add the following code to the `wire.go` file:

```go
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
```

Run `go mod tidy` again in the pongservice directory:

```bash
go mod tidy
```

Then execute the following wire command in the `pongservice` directory:

```bash
wire ./...
```

After executing the wire command, the `wire_gen.go` file will be automatically created in the `pongsevice/cmd/server`
directory.

### Generate main.go

The last step, create the main.go file in the `pongservice/cmd` directory

```go
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
```

## Run service

Execute the following command in the `pongservice` directory to start the microservice:

>
> **NOTICE**
> *Execute the command in the `pongservice` directory instead of the `pongservice/cmd` directory*
>

```bash
go run cmd/main.go
```

After starting the service, the following information will be displayed:

```text
Listening grpc server on port:50051
Listening http server on port:9001
```

Enter the following address in the browser to access the service:

```text
127.0.01:9001/pong.v1.pong?msg=best practice
```

In case of success, the following data will be returned:

```text
{
    "msg": "response pong msg:best practice"
}
```

# Congratulations

Now, we have learned how to create a project structure that can develop basic functional microservices.

In the next part, we continue to create a microservice called pingservice and access the pongservice we created in this
part.




