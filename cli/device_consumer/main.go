package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/sakalouski-alex/device-service/src/proto-gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var address string
	var help bool
	flag.StringVar(&address, "address", "localhost:50051", "The server address")
	flag.BoolVar(&help, "help", help, "Display help")
	flag.BoolVar(&help, "h", help, "Display help")
	flag.Parse()

	if help {
		log.Println(`Usage: main.go [OPTIONS] COMMAND

Options:
  -address string
        The server address (default "localhost:50051")

Commands:
  list    List all devices
  add     Add a new device
  delete  Delete a device by id
`)
		return
	}

	command := flag.Arg(0)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDeviceProviderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch command {
	case "add":
		fallthrough
	case "delete":
		log.Printf("Command not implemented: %s", command)
		return
	case "list":
		resp, err := client.ListDevices(ctx, &pb.GetDeviceListRequest{})
		if err != nil {
			log.Fatalf("could not list devices: %v", err)
		}

		for _, device := range resp.GetDevices() {
			log.Printf("Device: %v", device)
		}
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
