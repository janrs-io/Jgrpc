# Jgrpc

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
127.0.0.1:9002/ping.v1.ping?msg=best%20practice
```

It will return `json` data:

```json
{
    "msg": "response ping msg:best practice and msg from pong service is: response pong msg:request from ping service"
}
```