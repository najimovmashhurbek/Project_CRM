package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pb "github.com/najimovmashhurbek/crm_projekt/user-service.crm/genproto"
	//"google.golang.org/grpc/internal/status"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *pb.User) (*pb.User, error) {
	userr := pb.User{}
	query := "insert into users (id,company_id,firstname,lastname,email,password,phonenumber,role,gender,info,createdat) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id,company_id,firstname,lastname,email,password,phonenumber,role,gender,info,createdat"
	time := time.Now()
	id := uuid.New()
	id1 := uuid.New()
	err := r.db.QueryRow(query, id, id1, user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNumber, user.Role, user.Gender, user.Info, time).Scan(
		&userr.Id,
		&userr.CompanyId,
		&userr.FirstName,
		&userr.LastName,
		&userr.Email,
		&userr.Password,
		&userr.PhoneNumber,
		&userr.Role,
		&userr.Gender,
		&userr.Info,
		&userr.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &userr, nil
}
