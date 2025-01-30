package services

import (
	"context"
	"time"
)

type IService interface {
	Healthcheck(ctx context.Context) error

	RecordSignup(ctx context.Context, param *ParamRecordSignup) (*RespRecordSignup, error) // RecordSignup is the service to record the signup done by the user via 3rd party auth service
	DeleteAccount(ctx context.Context, param *ParamRequestAccountDeletion) error
}

// ParamSignup is the parameter for Signup
// No password is required, assuming we're using 3rd party auth service
// And we're just storing the email and phone number in our database (more secure)
type ParamRecordSignup struct {
	Email          *string `json:"email"`
	Phone          *string `json:"phone"`
	AuthProvider   *string `json:"auth_provider"`
	AuthProviderID *string `json:"auth_provider_id"`
}

type RespRecordSignup struct {
	Data *Signup `json:"data"`
}

type Signup struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Email          *string   `json:"email"`
	Phone          *string   `json:"phone"`
	AuthProvider   *string   `json:"auth_provider"`
	AuthProviderID *string   `json:"auth_provider_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ParamRequestAccountDeletion struct {
	ID string `json:"id"`
}

type AccountDeletion struct {
	DeletionDate time.Time `json:"deletion_date"`
}
