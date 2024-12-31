package pg

import (
	"log"
	"time"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/google/uuid"
)

type createAccount struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Email     *string   `db:"email"`
	Phone     *string   `db:"phone"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func mapParamCreateAccount(param *db.ParamCreateAccount) createAccount {
	return createAccount{
		ID: func() string {
			log.Println("ID: ", param.ID)
			if param.ID != nil {
				return *param.ID
			}
			return uuid.New().String()
		}(),
		Username: param.Username,
		Email:    param.Email,
		Phone:    param.Phone,
		Status: func() string {
			if param.Status != nil {
				return *param.Status
			}
			return "active"
		}(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
