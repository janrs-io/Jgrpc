# [Jgrpc](https://github.com/janrs-io/Jgrpc "Jgrpc")

This [project](https://github.com/janrs-io/Jgrpc "project") provides a reference for best practices for developing
microservices based on `Go/Grpc/kubernetes/Istio`.

And it implements `CICD` based on `Jenkins/Gitlab/Harbor`.

And use the [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway "grpc-gateway") as a gateway proxy.

This best practice is divided into three parts:

1. Create a microservice named `pongservice`
2. Create a microservice named `pingservice` and access `pongservice`
3. Create a `CICD` deployment process base on `Jenkins/Gitlab/Harbor`

# QuickStart

Just clone the project and run `go run cmd/main.go` in the `pingservice` and `pongservice` directory.

And enter this address in browser to visit `pong` http server:

```text
http://127.0.0.1:9001/pong.v1.pong?msg=best%20practive
```

It will return `json` data:

```json
{
    "msg": "response pong msg:best practive"
}
```

And enter this address in browser to visit `pingservice` http server:

```text
http://127.0.0.1:9002/ping.v1.ping?msg=best%20practice
```

It will return `json` data:

```json
{
    "msg": "response ping msg:best practice and msg from pong service is: response pong msg:request from ping service"
}
```