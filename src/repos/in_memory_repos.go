package repos

import (
	pb "github.com/sakalouski-alex/device-service/src/proto-gen/proto"
)

type MemoryDeviceRepo struct {
	devices []*pb.Device
}

func NewMemoryDeviceRepo(devices []*pb.Device) *MemoryDeviceRepo {
	return &MemoryDeviceRepo{devices: devices}
}

func (repo *MemoryDeviceRepo) ListDevices() ([]*pb.Device, error) {
	return repo.devices, nil
}
