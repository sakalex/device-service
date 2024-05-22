package main

import (
	"flag"

	device_service "github.com/sakalouski-alex/device-service/src/devices"
	"github.com/sakalouski-alex/device-service/src/repos"

	pb "github.com/sakalouski-alex/device-service/src/proto-gen/proto"
)

func main() {
	var port int = 50051
	flag.IntVar(&port, "port", port, "The server port")
	flag.Parse()

	devices := []*pb.Device{
		{Id: "mouse_1", Type: pb.Device_MOUSE, DevicePath: "/dev/input/mouse0", VendorId: "uuid1", ProductId: "uuid2"},
		{Id: "keyboard_1", Type: pb.Device_KEYBOARD, DevicePath: "/dev/input/keyboard1", VendorId: "uuid3", ProductId: "uuid4"},
	}

	repo := repos.NewMemoryDeviceRepo(devices)
	provider := device_service.NewDeviceProvider(repo)

	config := device_service.DeviceServiceConfig{
		Port:     port,
		Provider: provider,
	}

	device_service.Start(config)
}
