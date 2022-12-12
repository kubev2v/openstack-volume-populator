FROM golang:1.19.3-alpine3.16
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY ./pkg ./pkg
COPY main.go ./

RUN go mod download
RUN go build -o /main

ENTRYPOINT ["/main"]
