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

func UT_AccountRepository_FindAccountByID(t *testing.T) {
	// instantiate the mocked repository
	accRepo, _, readMock, err := newAccountRepositoryWithMock()
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
				expAccAv := accToArrDV(exp)

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
