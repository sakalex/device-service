package repos

import (
	"fmt"
	"sync"

	pb "github.com/sakalouski-alex/device-service/src/proto-gen/proto"
)

type MemoryDeviceRepo struct {
	mu      sync.Mutex
	devices map[string]*pb.Device
}

func NewMemoryDeviceRepo(devices []*pb.Device) *MemoryDeviceRepo {
	deviceMap := make(map[string]*pb.Device)
	for _, device := range devices {
		deviceMap[device.Id] = device
	}
	return &MemoryDeviceRepo{devices: deviceMap}
}

func (r *MemoryDeviceRepo) ListDevices() ([]*pb.Device, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	devices := make([]*pb.Device, 0, len(r.devices))
	for _, device := range r.devices {
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *MemoryDeviceRepo) AddDevice(device *pb.Device) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.devices[device.Id]; exists {
		return fmt.Errorf("device with id=%s already exists", device.Id)
	}

	r.devices[device.Id] = device
	return nil
}

func (r *MemoryDeviceRepo) DeleteDevice(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.devices[id]; !exists {
		return fmt.Errorf("device with id=%s not found", id)
	}

	delete(r.devices, id)
	return nil
}
