package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/najimovmashhurbek/crm_projekt/companies-service.crm/storage/postgres"
	"github.com/najimovmashhurbek/crm_projekt/companies-service.crm/storage/repo"
)

//IStorage ...
type IStorage interface {
	Companies() repo.CompaniesStorageI
}

type storagePg struct {
	db            *sqlx.DB
	companiesRepo repo.CompaniesStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:            db,
		companiesRepo: postgres.NewCompaniesRepo(db),
	}
}

func (s storagePg) Companies() repo.CompaniesStorageI {
	return s.companiesRepo
}
