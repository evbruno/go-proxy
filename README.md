## Proxy

**Use this proxy for local development of your Docker containers**

## Running

1st: start the service you want to proxy, eg:

```
docker run --rm --name my-service -p 8000:8000 educhaos/nodejs-server-sandbox:latest
```

2nd: start the proxy, linking the service you've started:

```
docker run --rm --name my-proxy -p 9090:9090 \
  --link my-service:my-service \
  educhaos/go-proxy:latest \
  go run proxy.go -url http://my-service:8000/
```

3rd: test you proxy:

```
curl -v localhost:9090/ping
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 9090 (#0)
> GET /ping HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Length: 4
< Content-Type: text/html; charset=utf-8
< Date: Mon, 06 Dec 2021 14:52:43 GMT
< Etag: W/"4-DlFKBmK8tp3IY5U9HOJuPUDoGoc"
< X-Goproxy: GoProxy
< X-Powered-By: Express
<
* Connection #0 to host localhost left intact
pong* Closing connection 0
```

## Configs

```
-port     (default to 9090)
-url      (default to http://127.0.0.1:8080)
-timeout  (default to 0)
```