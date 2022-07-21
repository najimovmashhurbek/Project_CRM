package grpcClient

import (
	"fmt"

	"github.com/najimovmashhurbek/crm_projekt/user-service.crm/config"
	pb "github.com/najimovmashhurbek/crm_projekt/user-service.crm/genproto"
	"google.golang.org/grpc"
)

//GrpcClientI ...
type GrpcClientI interface {
	CompaniesService() pb.CompaniesServiceClient
	SheduleService() pb.ScheduleServiceClient
}

//GrpcClient ...
type GrpcClient struct {
	cfg              config.Config
	companiesService pb.CompaniesServiceClient
	scheduleService pb.ScheduleServiceClient
}

//New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connComp, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CompaniesServiceHost, cfg.CompaniesServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d",
			cfg.CompaniesServiceHost, cfg.CompaniesServicePort)
	}

	connSchedule, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ScheduleServiceHost, cfg.ScheduleServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("schedule service dial host: %s port: %d",
			cfg.ScheduleServiceHost, cfg.ScheduleServicePort)
	}

	return &GrpcClient{
		cfg:              cfg,
		companiesService: pb.NewCompaniesServiceClient(connComp),
		scheduleService: pb.NewScheduleServiceClient(connSchedule),
	}, nil
}



func (s *GrpcClient) CompaniesService() pb.CompaniesServiceClient {
	return s.companiesService
}
func (s *GrpcClient) SheduleService() pb.ScheduleServiceClient {
	return s.scheduleService
}
