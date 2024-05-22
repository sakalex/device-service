FROM golang:1.19

RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    make

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

COPY . .
