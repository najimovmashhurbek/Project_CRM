package grpcClient

import (
	"fmt"

	"github.com/najimovmashhurbek/crm_projekt/companies-service.crm/config"
	pb "github.com/najimovmashhurbek/crm_projekt/companies-service.crm/genproto"
	"google.golang.org/grpc"
)

//GrpcClientI ...
type GrpcClientI interface {
	UserService() pb.UserServiceClient
	Schedule() pb.ScheduleServiceClient
}

//GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	userService pb.UserServiceClient
	scheduleService pb.ScheduleServiceClient
}

//New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port: %d",
			cfg.UserServiceHost, cfg.UserServicePort)
	}

	connSchedule, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ScheduleServiceHost, cfg.ScheduleServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("schedule service dial host: %s port: %d",
			cfg.ScheduleServiceHost, cfg.ScheduleServicePort)
	}

	return &GrpcClient{
		cfg:         cfg,
		userService: pb.NewUserServiceClient(connUser),
		scheduleService: pb.NewScheduleServiceClient(connSchedule),
	}, nil
}




func (s *GrpcClient) UserService() pb.UserServiceClient {
	return s.userService
}
func (s *GrpcClient) Schedule() pb.ScheduleServiceClient {
	return s.scheduleService
}