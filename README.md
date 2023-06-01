# Jgrpc

**如果觉得该项目对你的学习有帮助，点个 star**

## 完整的案例代码直达

https://github.com/janrs-io/Jgrpc-example

中文文档地址：

###### [基于Go/Grpc/kubernetes/Istio开发微服务的最佳实践尝试 - 1/3](https://janrs.com/br6f "基于Go/Grpc/kubernetes/Istio开发微服务的最佳实践尝试 - 1/3")

###### [基于Go/Grpc/kubernetes/Istio开发微服务的最佳实践尝试 - 2/3](https://janrs.com/ugj7 "基于Go/Grpc/kubernetes/Istio开发微服务的最佳实践尝试 - 2/3")

###### [基于Go/Grpc/kubernetes/Istio开发微服务的最佳实践尝试 - 3/3](https://janrs.com/6rdh "基于Go/Grpc/kubernetes/Istio开发微服务的最佳实践尝试 - 3/3")

This project provides a reference for best practices for developing
microservices based on `Go/Grpc/kubernetes/Istio`.

And it implements `CICD` based on `Jenkins/Gitlab/Harbor`.

And use the [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway "grpc-gateway") as a gateway proxy.

This best practice is divided into three parts:

1. [Create a microservice named `pongservice`](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part01.md "Create a microservice named `pongservice`")
2. [Create a microservice named `pingservice` and access `pongservice`](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part02.md "Create a microservice named `pingservice` and access `pongservice`")
3. [Create a `CICD` deployment process base on `Jenkins/Gitlab/Harbor`](https://github.com/janrs-io/Jgrpc/blob/master/best-practices-for-developing-microservices-based-on-go-grpc-kubernetes-Istio-part03.md "Create a `CICD` deployment process base on `Jenkins/Gitlab/Harbor`")

# QuickStart

Just clone the project and run `go run cmd/main.go` in the `pingservice` and `pongservice` directory.

And enter this address in browser to visit `pong` http server:

```text
http://127.0.0.1:9001/pong.v1.pong?msg=best%20practice
```

It will return `json` data:

```json
{
    "msg": "response pong msg:best practice"
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