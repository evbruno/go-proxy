FROM golang:1.17-alpine

WORKDIR /app

COPY *.go ./

EXPOSE 9090

CMD go run proxy.go