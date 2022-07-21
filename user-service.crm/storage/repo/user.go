package repo

import (
	pb "github.com/najimovmashhurbek/crm_projekt/user-service.crm/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User, error)
}
