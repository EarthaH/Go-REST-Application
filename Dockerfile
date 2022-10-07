# syntax=docker/dockerfile:1

FROM golang:1.19-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd/main/*.go ./cmd/main/
COPY pkg/api/*.go ./pkg/api/
COPY pkg/dto/*.go ./pkg/dto/
COPY pkg/logger/*.go ./pkg/logger/
COPY pkg/ls/*.go ./pkg/ls/

RUN go build ./cmd/main/main.go

EXPOSE 8080

CMD [ "./main" ]
