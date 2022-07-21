package postgres

import (
	//"time"

	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	//"github.com/lib/pq"
	pb "github.com/najimovmashhurbek/crm_projekt/companies-service.crm/genproto"
	//"google.golang.org/grpc/internal/status"
)

type companiesRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewCompaniesRepo(db *sqlx.DB) *companiesRepo {
	return &companiesRepo{db: db}
}

func (r *companiesRepo) CreateCompanies(companies *pb.Companies) (*pb.Companies, error) {
	companiess := pb.Companies{}
	query := "insert into companies (id,name,email,password,phonenumber,info,createdat) values ($1,$2,$3,$4,$5,$6,$7) RETURNING id,name,email,password,phonenumber,info,createdat"
	id := uuid.New()
	time := time.Now()
	err := r.db.QueryRow(query, id, companies.Name, companies.Email, companies.Password, companies.PhoneNumbers, companies.Info, time).Scan(
		&companiess.Id,
		&companiess.Name,
		&companiess.Email,
		&companiess.Password,
		&companiess.PhoneNumbers,
		&companiess.Info,
		&companiess.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &companiess, nil
}
