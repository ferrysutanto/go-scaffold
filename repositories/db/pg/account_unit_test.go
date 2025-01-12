package pg_test

import (
	"context"
	"testing"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/ferrysutanto/go-scaffold/repositories/db/pg"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func UT_AccountRepository_CreateAccount(t *testing.T) {
	// instantiate the mocked repository every time the test runs
	accRepo, writeMock, _, err := newMockAccountRepository()
	if err != nil {
		require.NoError(t, err)
	}

	type args struct {
		ctx   context.Context
		param *db.ParamCreateAccount
	}
	tests := []struct {
		name    string
		setup   func()
		args    args
		want    *db.Account
		wantErr error
	}{
		// Positive test cases
		{
			name: "positive - success",
			setup: func() {
				exp := accList[0]
				arg := dbAccToArgCrtAccADV(exp)
				expAccAv := dbAccToADV(exp)

				writeMock.ExpectQuery(accQueryInsert).
					WithArgs(arg...).
					WillReturnRows(
						writeMock.NewRows(accColumns).
							AddRow(expAccAv...))
			},
			args:    args{ctx: context.Background(), param: &db.ParamCreateAccount{ID: &accList[0].ID, Username: accList[0].Username, Email: accList[0].Email, Phone: accList[0].Phone, Status: &accList[0].Status}},
			want:    &db.Account{ID: accList[0].ID, Username: accList[0].Username, Email: accList[0].Email, Phone: accList[0].Phone, Status: accList[0].Status, CreatedAt: accList[0].CreatedAt, UpdatedAt: accList[0].UpdatedAt},
			wantErr: nil,
		},
		// Negative test cases
		{name: "negative 1 - empty username", setup: func() {}, args: args{ctx: context.Background(), param: &db.ParamCreateAccount{Username: ""}}, want: nil, wantErr: pg.ErrValidationFailed([]string{pg.ErrUsernameRequired.Error(), pg.ErrEmailOrPhoneRequired.Error()})},
		{name: "negative 2 - empty email and phone", setup: func() {}, args: args{ctx: context.Background(), param: &db.ParamCreateAccount{Username: faker.Username()}}, want: nil, wantErr: pg.ErrValidationFailed([]string{pg.ErrEmailOrPhoneRequired.Error()})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := accRepo.CreateAccount(tt.args.ctx, tt.args.param)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func UT_AccountRepository_FindAccountByID(t *testing.T) {
	// instantiate the mocked repository
	accRepo, _, readMock, err := newMockAccountRepository()
	if err != nil {
		require.NoError(t, err)
	}

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		setup   func()
		args    args
		want    *db.Account
		wantErr error
	}{
		// Positive test cases
		{
			name: "positive - success",
			setup: func() {
				exp := accList[0]
				expAccAv := dbAccToADV(exp)

				readMock.ExpectQuery(accQueryFindByID).
					WithArgs(exp.ID).
					WillReturnRows(
						readMock.NewRows(accColumns).
							AddRow(expAccAv...))
			},
			args:    args{ctx: context.Background(), id: accList[0].ID},
			want:    accList[0],
			wantErr: nil,
		},
		// Negative test cases
		{name: "negative 1 - empty id supplied", setup: func() {}, args: args{ctx: context.Background(), id: ""}, want: nil, wantErr: pg.ErrIdRequired},
		{name: "negative 2 - invalid id", setup: func() {}, args: args{ctx: context.Background(), id: faker.Sentence()}, want: nil, wantErr: pg.ErrInvalidUUID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := accRepo.FindAccountByID(tt.args.ctx, tt.args.id)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func UT_AccountRepository_UpdateAccount(t *testing.T) {
	// instantiate the mocked repository
	accRepo, writeMock, _, err := newMockAccountRepository()
	if err != nil {
		require.NoError(t, err)
	}

	type args struct {
		ctx   context.Context
		param *db.ParamUpdateAccount
	}
	tests := []struct {
		name    string
		setup   func()
		args    args
		want    *db.Account
		wantErr error
	}{
		// Positive test cases
		{
			name: "positive - success",
			setup: func() {
				exp := accList[0]
				expAccAv := dbAccToADV(exp)

				writeMock.ExpectQuery(accQueryUpdate).
					WithArgs(dbAccToArgUpdAccADV(exp)...).
					WillReturnRows(
						writeMock.NewRows(accColumns).
							AddRow(expAccAv...))
			},
			args: args{
				ctx: context.Background(),
				param: &db.ParamUpdateAccount{
					ID: accList[0].ID, Username: accList[0].Username, Email: accList[0].Email, Phone: accList[0].Phone, Status: accList[0].Status}},
			want:    accList[0],
			wantErr: nil,
		},
		// Negative test cases
		// {name: "negative 1 - empty id supplied", setup: func() {}, args: args{ctx: context.Background(), id: ""}, want: nil, wantErr: pg.ErrIdRequired},
		// {name: "negative 2 - invalid id", setup: func() {}, args: args{ctx: context.Background(), id: faker.Sentence()}, want: nil, wantErr: pg.ErrInvalidUUID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := accRepo.UpdateAccount(tt.args.ctx, tt.args.param)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
