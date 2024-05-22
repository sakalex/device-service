package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/sakalouski-alex/device-service/src/proto-gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
)

type DeviceType struct {
	value *pb.Device_Type
}

func (dt *DeviceType) String() string {
	if dt.value == nil {
		return ""
	}
	return dt.value.String()
}

func (dt *DeviceType) Set(s string) error {
	val, ok := pb.Device_Type_value[strings.ToUpper(s)]

	if !ok {
		return fmt.Errorf("invalid device type: \"%s\"", s)
	}

	dt.value = pb.Device_Type(val).Enum()
	return nil
}

func getEnumValues(enum map[string]int32) []string {
	values := make([]string, 0, len(enum))
	for key := range enum {
		values = append(values, key)
	}
	return values
}

func main() {
	var address string
	var help bool
	var id string
	var deviceType = DeviceType{value: pb.Device_MOUSE.Enum()}
	var path string
	var vendorId string
	var productId string

	flag.StringVar(&address, "address", "localhost:50051", "The server address")

	listFlag := flag.NewFlagSet("list", flag.ExitOnError)
	deleteFlags := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteFlags.StringVar(&id, "id", uuid.New().String(), "The device id")

	addFlags := flag.NewFlagSet("add", flag.ExitOnError)
	addFlags.StringVar(&id, "id", uuid.New().String(), "The device id")
	addFlags.Var(&deviceType, "type", fmt.Sprintf("The device type (%s). Default is MOUSE.", strings.Join(getEnumValues(pb.Device_Type_value), ", ")))
	addFlags.StringVar(&path, "path", "", "The device path")
	addFlags.StringVar(&vendorId, "vendor_id", uuid.New().String(), "The vendor id")
	addFlags.StringVar(&productId, "product_id", uuid.New().String(), "The product id")

	flag.BoolVar(&help, "help", help, "Display help")
	flag.Parse()

	if flag.NArg() == 0 || help {
		flag.Usage()
		fmt.Println(`
Commands:
  list    List all devices
  add     Add a new device
  delete  Delete a device
`)
		listFlag.Usage()
		addFlags.Usage()
		deleteFlags.Usage()
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
		for i, v := range os.Args {
			if v == "add" {
				if err := addFlags.Parse(os.Args[i+1:]); err != nil {
					flag.Usage()
					addFlags.Usage()
					return
				}
				break
			}
		}
		_, err := client.AddDevice(ctx, &pb.AddDeviceRequest{
			Device: &pb.Device{
				Id:         id,
				Type:       *deviceType.value,
				DevicePath: path,
				VendorId:   vendorId,
				ProductId:  productId,
			},
		})
		if err != nil {
			log.Fatalf("could not add device: %v", err)
		}
	case "delete":
		for i, v := range os.Args {
			if v == "delete" {
				if err := deleteFlags.Parse(os.Args[i+1:]); err != nil {
					flag.Usage()
					deleteFlags.Usage()
					return
				}
				break
			}
		}
		_, err := client.DeleteDevice(ctx, &pb.DeleteDeviceRequest{Id: id})
		if err != nil {
			log.Fatalf("could not delete device: %v", err)
		}
	case "list":
		for i, v := range os.Args {
			if v == "list" {
				if err := listFlag.Parse(os.Args[i+1:]); err != nil {
					flag.Usage()
					listFlag.Usage()
					return
				}
				break
			}
		}
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
