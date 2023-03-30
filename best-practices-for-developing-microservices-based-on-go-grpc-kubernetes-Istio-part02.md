In the previous part, we created the directory structure of the microservice and implemented the
microservice `pongservice`.

In this part, we continue to implement a microservice name `pingservice`, and access the `pongservice` microservice that
has been deployed in the previous part.

Creating a new microservice is very simple, just copy the `pongservice` microservice that has been created before, and
then make some small changes.

# Project Structure

The final structure of the directory in this part is as follows:

```bash
pingservice
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
│   ├── client.go
│   ├── config.go
│   └── config.yaml
├── genproto
│   └── v1
│       ├── gw
│       │   └── pingservice.pb.gw.go
│       ├── pingservice.pb.go
│       └── pingservice_grpc.pb.go
├── go.mod
├── go.sum
├── proto
│   ├── buf.lock
│   ├── buf.yaml
│   └── v1
│       ├── pingservice.proto
│       └── pingservice.yaml
└── service
    ├── client.go
    └── server.go

9 directories, 21 files
```

# Start

## Copy

Execute the following copy command in the src directory:

```bash
cp -R pongservice pingservice
```

## Delete wire_gen.go

Delete `wire_gen.go` file in the `pingservice/cmd/server` directory.

## Modify go.mod module

Modify the module of the `go.mod` file as shown in the following code:

```go
module github.com/janrs-io/Jgrpc/src/pingservice
```

## Regenerate proto

Delete the `pongservice.proto` and `pongservice.yaml` files in the `pingservice/proto/v1` directory and the
entire `genproto` folder.

Then recreate the `pingservice.proto` and `pingservice.yaml` files, the code is as follows:

`pingservice.proto` code:

```proto
syntax = "proto3";

package proto.v1;

option go_package = "github.com/janrs-io/Jgrpc/src/pingservice/genproto/v1";

service PingService {
  rpc Ping(PingRequest) returns(PingResponse){}
}

message PingRequest {
  string msg = 1 ;
}

message PingResponse {
  string msg = 1;
}
```

`pingservice.yaml` code:

```yaml
type: google.api.Service
config_version: 3

http:
  rules:
    - selector: proto.v1.PingService.Ping
      get: /ping.v1.ping
```

Modify the `buf.gen.yaml` file in the `pingservice` directory. The complete code after modification is as follows:

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
      - grpc_api_configuration=proto/v1/pingservice.yaml
      - standalone=true
```

Execute the following command in the `pingservice` directory to generate the proto file:

```bash
buf generate proto/v1
```

After executing the command, the `genproto` directory will be regenerated and the `*pb.go` file will be created
automatically.

## Modify import path & codes

Check the import of **all files**, and change the `pongservice` of the import path to `pingservice`.

Change the `Pong/pong` of all codes to `Ping/ping` until there are no errors.

## Modify config.yaml

Modify the `config.yaml` file in the config directory.
Change the server port of grpc to `50052` and the server port of http to `9002`.
Change the service name of grpc to `ping-grpc` and the service name of http to `ping-http`.

The complete code after modification is as follows:

```yaml
# grpc config
grpc:
  host: ""
  port: ":50052"
  name: "ping-grpc"

# http config
http:
  host: ""
  port: ":9002"
  name: "ping-http"
```

## Wire generate inject

Execute the following wire command in the `pingservice` directory to regenerate the dependency injection files:

```bash
wire ./...
```

## Import pongservier's go.mod

We want to access the grpc service of `pongservice` in this microservice of `pingservice`.So the `go.mod`
of `pongservice` needs to be imported.

Modify the `go.mod` file in the `pingservice` directory and add the code to import `pongservice`, as follows:

```go
module github.com/janrs-io/Jgrpc/src/pingservice

go 1.19

replace (
	pongservice => ../pongservice
)

require (
	pongservice v0.0.0
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/spf13/viper v1.15.0
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

// the other required package...
```

## Modify config.yaml

Modify the `config.yaml` file in the `pingservice/config` directory and add the following code:

```yaml
# service client
client:
  pong: ":50051"
```

The complete code after modification is as follows:

```yaml
# grpc config
grpc:
  host: ""
  port: ":50052"
  name: "ping-grpc"

# http config
http:
  host: ""
  port: ":9002"
  name: "ping-http"

# service client
client:
  pong: ":50051"
```

## Generate client.go

Generate `client.go` file in the `pingservice/config` directory and add the following code:

```go
package config

// Client Client service config
type Client struct {
	Pong string `json:"pong" yaml:"pong"`
}
```

## Modify config.go

Add a `Client` field in the `Config` struct as the following code:

```go
Client Client `json:"client" yaml:"client"`
```

The complete code after modification is as follows:

```go
package config

import (
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Config Service config
type Config struct {
	Grpc   Grpc   `json:"grpc" yaml:"grpc"`
	Http   Http   `json:"http" yaml:"http"`
	Client Client `json:"client" yaml:"client"`
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

## Modify client.go

Modify the `client.go` file under the `pingservice/service` directory and add the following code:

```go
// NewPongClient New pong service client
func NewPongClient(conf *config.Config) (pongclientv1.PongServiceClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, conf.Client.Pong, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("dial auth server failed.[ERROR]=>" + err.Error())
		return nil, err
	}
	client := pongclientv1.NewPongServiceClient(conn)
	return client, nil

}
```

The complete code after modification is as follows:

```go
package service

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/janrs-io/Jgrpc/src/pingservice/config"
	v1 "github.com/janrs-io/Jgrpc/src/pingservice/genproto/v1"
	pongclientv1 "github.com/janrs-io/Jgrpc/src/pongservice/genproto/v1"
)

// NewClient New service's client
func NewClient(conf *config.Config) (v1.PingServiceClient, error) {

	serverAddress := conf.Grpc.Host + conf.Grpc.Port
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := v1.NewPingServiceClient(conn)
	return client, nil

}

// NewPongClient New pong service client
func NewPongClient(conf *config.Config) (pongclientv1.PongServiceClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, conf.Client.Pong, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("dial auth server failed.[ERROR]=>" + err.Error())
		return nil, err
	}
	client := pongclientv1.NewPongServiceClient(conn)
	return client, nil

}
```

## Modify server.go

Modify the `server.go` file under the `pingservice/service` directory,the complete code after modification is as
follows:

```go
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
```

## Modify wire.go

Modify the `wire.go` file of `pingservice/cmd/server` and add `service.NewPongClient` dependency injection. code show as
below:

```go
service.NewPongClient
```

The complete code after modification is as follows:

```go
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
```

Execute the following wire command in the `pingservice` directory to regenerate the dependency injection file:

> *If the go.mod import errors shown up,just run `go mod tidy` again in the `pingservice` directory.*

```bash
wire ./...
```

## Run service

Execute the `go run` command in the `pongservice` directory and `pingservice` directory respectively.

```bash
go run cmd/main.go
```

Enter the following request address in the browser:

```text
127.0.01:9002/ping.v1.ping?msg=best practice
```

Everything correct returns the following json data:

```json
{
    "msg": "response ping msg:best practice and msg from pong service is: response pong msg:request from ping service"
}
```

# Congratulations

In this part, we create a new microservice of `pingservice` and implement the grpc service of accessing `pongservice`.

I believe that through these two simple attempts to create microservices, you must feel that it is not difficult to
develop microservices based on `Go` and `Grpc`.

In the next part, we will leverage `Jenkins/Gitlab/Harbor` and `Kubernets/Istio` for `CICD` deployment for `devops`.








