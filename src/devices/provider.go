package device

import (
	"context"

	pb "github.com/sakalouski-alex/device-service/src/proto-gen/proto"
	"github.com/sakalouski-alex/device-service/src/repos"
)

type DeviceProvider struct {
	pb.UnimplementedDeviceProviderServiceServer
	repo repos.DeviceRepo
}

func NewDeviceProvider(repo repos.DeviceRepo) pb.DeviceProviderServiceServer {
	return &DeviceProvider{repo: repo}
}

func (s *DeviceProvider) ListDevices(ctx context.Context, in *pb.GetDeviceListRequest) (*pb.DeviceListResponse, error) {
	devices, err := s.repo.ListDevices()
	if err != nil {
		return nil, err
	}
	return &pb.DeviceListResponse{Devices: devices}, nil
}

func (s *DeviceProvider) AddDevice(ctx context.Context, in *pb.AddDeviceRequest) (*pb.OperationStatus, error) {
	err := s.repo.AddDevice(in.Device)
	if err != nil {
		return &pb.OperationStatus{Success: false}, err
	}
	return &pb.OperationStatus{Success: true}, nil
}

func (s *DeviceProvider) DeleteDevice(ctx context.Context, in *pb.DeleteDeviceRequest) (*pb.OperationStatus, error) {
	err := s.repo.DeleteDevice(in.Id)
	if err != nil {
		return &pb.OperationStatus{Success: false}, err
	}
	return &pb.OperationStatus{Success: true}, nil
}
