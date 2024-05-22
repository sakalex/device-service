package repos

import (
	pb "github.com/sakalouski-alex/device-service/src/proto-gen/proto"
)

type DeviceRepo interface {
	ListDevices() ([]*pb.Device, error)
}
