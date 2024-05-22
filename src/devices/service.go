package device

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/golang/glog"
	pb "github.com/sakalouski-alex/device-service/src/proto-gen/proto"

	"google.golang.org/grpc"
)

type DeviceServiceConfig struct {
	Port     int
	Provider pb.DeviceProviderServiceServer
}

func Start(config DeviceServiceConfig) {
	glog.Infof("Starting the server on port %d...", config.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		glog.Fatalf("Failed to listen on port %d: %v", config.Port, err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			start := time.Now()
			resp, err := handler(ctx, req)
			if err != nil {
				glog.Errorf("gRPC Request: %s, Time: %v, Error: %v", info.FullMethod, time.Since(start), err)
			} else {
				glog.Infof("gRPC Request: %s, Time: %v, Response: %v", info.FullMethod, time.Since(start), resp)
			}
			return resp, err
		}),
	)
	pb.RegisterDeviceProviderServiceServer(s, config.Provider)

	glog.Infof("Server is running on port %d", config.Port)

	if err := s.Serve(lis); err != nil {
		glog.Fatalf("Failed to serve: %v", err)
	}
}
