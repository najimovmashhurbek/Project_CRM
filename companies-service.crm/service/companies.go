package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/najimovmashhurbek/crm_projekt/companies-service.crm/genproto"
	l "github.com/najimovmashhurbek/crm_projekt/companies-service.crm/pkg/logger"
	cl "github.com/najimovmashhurbek/crm_projekt/companies-service.crm/service/grpc_client"
	"github.com/najimovmashhurbek/crm_projekt/companies-service.crm/storage"
)

//UserService ...
type CompaniesService struct {
	storage storage.IStorage
	logger  l.Logger
	client  cl.GrpcClientI
}

//NewUserService ...
func NewCompaniesService(db *sqlx.DB, log l.Logger, client cl.GrpcClientI) *CompaniesService {
	return &CompaniesService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}
func (s *CompaniesService) CreateCompanies(ctx context.Context, req *pb.Companies) (*pb.Companies, error) {
	comp, err := s.storage.Companies().CreateCompanies(req)
	if err != nil {
		return nil, err
	}
	if req.User != nil {
		for _, usr := range req.User {
			users, err := s.client.UserService().CreateUser(ctx, usr)
			if err != nil {
				s.logger.Error("failed while inserting user post", l.Error(err))
				return nil, err
			}
			comp.User = append(comp.User, users)
		}
	}
	return comp, err
}
