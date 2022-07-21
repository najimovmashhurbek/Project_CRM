package repo

import (
	pb "github.com/najimovmashhurbek/crm_projekt/companies-service.crm/genproto"
)

//UserStorageI ...
type CompaniesStorageI interface {
	CreateCompanies(*pb.Companies) (*pb.Companies, error)
}
