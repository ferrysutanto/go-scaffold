package pg

import (
	"time"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

type profileEntity struct {
	ID         string     `db:"id"`
	AccountID  string     `db:"account_id"`
	FirstName  string     `db:"first_name"`
	LastName   string     `db:"last_name"`
	Birthdate  *time.Time `db:"birthdate,omitempty"`
	Sex        *string    `db:"sex,omitempty"` // biological sex of the individual.
	PictureURL *string    `db:"picture_url,omitempty"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
}

func mapProfileEntityToProfile(entity *profileEntity) *db.Profile {
	return &db.Profile{
		ID:         entity.ID,
		AccountID:  entity.AccountID,
		FirstName:  entity.FirstName,
		LastName:   entity.LastName,
		Birthdate:  entity.Birthdate,
		Sex:        entity.Sex,
		PictureURL: entity.PictureURL,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}

type argCreateProfile struct {
	ID         string     `db:"id"`
	AccountID  string     `db:"account_id"`
	FirstName  string     `db:"first_name"`
	LastName   string     `db:"last_name"`
	Birthdate  *time.Time `db:"birthdate,omitempty"`
	Sex        *string    `db:"sex,omitempty"` // biological sex of the individual.
	PictureURL *string    `db:"picture_url,omitempty"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
}

func mapParamCreateProfile(param *db.ParamCreateProfile) argCreateProfile {
	return argCreateProfile{
		ID:         param.ID,
		AccountID:  param.AccountID,
		FirstName:  param.FirstName,
		LastName:   param.LastName,
		Birthdate:  param.Birthdate,
		Sex:        param.Sex,
		PictureURL: param.PictureURL,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type argUpdateProfile struct {
	ID         string     `db:"id"`
	AccountID  string     `db:"account_id"`
	FirstName  string     `db:"first_name"`
	LastName   string     `db:"last_name"`
	Birthdate  *time.Time `db:"birthdate,omitempty"`
	Sex        *string    `db:"sex,omitempty"` // biological sex of the individual.
	PictureURL *string    `db:"picture_url,omitempty"`
	UpdatedAt  time.Time  `db:"updated_at"`
}

func mapParamUpdateProfile(param *db.ParamUpdateProfile) argUpdateProfile {
	return argUpdateProfile{
		ID:         param.ID,
		AccountID:  param.AccountID,
		FirstName:  param.FirstName,
		LastName:   param.LastName,
		Birthdate:  param.Birthdate,
		Sex:        param.Sex,
		PictureURL: param.PictureURL,
		UpdatedAt:  time.Now(),
	}
}
