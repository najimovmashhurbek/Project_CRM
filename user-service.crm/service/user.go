package service

import (
	"context"
	_ "errors"

	"github.com/jmoiron/sqlx"
	pb "github.com/najimovmashhurbek/crm_projekt/user-service.crm/genproto"
	l "github.com/najimovmashhurbek/crm_projekt/user-service.crm/pkg/logger"
	"github.com/najimovmashhurbek/crm_projekt/user-service.crm/storage"

	//"golang.org/x/vuln/client"
	cl "github.com/najimovmashhurbek/crm_projekt/user-service.crm/service/grpc_client"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  cl.GrpcClientI
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client cl.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
