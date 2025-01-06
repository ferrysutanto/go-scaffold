package pg

import (
	"log"
	"time"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/google/uuid"
)

type accountEntity struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Email     *string   `db:"email"`
	Phone     *string   `db:"phone"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func mapAccountEntityToAccount(entity *accountEntity) *db.Account {
	return &db.Account{
		ID:        entity.ID,
		Username:  entity.Username,
		Email:     entity.Email,
		Phone:     entity.Phone,
		Status:    entity.Status,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

type argCreateAccount struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Email     *string   `db:"email"`
	Phone     *string   `db:"phone"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func mapParamCreateAccount(param *db.ParamCreateAccount) argCreateAccount {
	return argCreateAccount{
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

type argUpdateAccount struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Email     *string   `db:"email,omitempty"`
	Phone     *string   `db:"phone,omitempty"`
	Status    string    `db:"status"`
	UpdatedAt time.Time `db:"updated_at"`
}

func mapParamUpdateAccount(param *db.ParamUpdateAccount) argUpdateAccount {
	return argUpdateAccount{
		ID:        param.ID,
		Username:  param.Username,
		Email:     param.Email,
		Phone:     param.Phone,
		Status:    param.Status,
		UpdatedAt: time.Now(),
	}
}
