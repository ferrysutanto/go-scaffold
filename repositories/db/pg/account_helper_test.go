package pg_test

import (
	"context"
	"database/sql/driver"
	"time"

	"math/rand"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/ferrysutanto/go-scaffold/repositories/db/pg"
	"github.com/go-faker/faker/v4"
)

const (
	accQueryFindByID = "SELECT id, username, email, phone, status, created_at, updated_at FROM accounts WHERE"
	accQueryInsert   = `INSERT INTO accounts \(id, username, email, phone, status, created_at, updated_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\) RETURNING id, username, email, phone, status, created_at, updated_at`
	accQueryUpdate   = `UPDATE accounts SET username = \$1, email = \$2, phone = \$3, status = \$4 WHERE id = \$5 RETURNING id, username, email, phone, status, created_at, updated_at`
	accQueryDelete   = `DELETE FROM accounts WHERE id = \$1`
)

var (
	accColumns = []string{"id", "username", "email", "phone", "status", "created_at", "updated_at"}
	accStatus  = []string{"active", "inactive"}
)

// newMockAccountRepository creates a new AccountRepository with mock DB Conns (write and read)
func newMockAccountRepository() (resp db.IAccountRepository, writeMock sqlmock.Sqlmock, readMock sqlmock.Sqlmock, err error) {
	write, writeMock, err := sqlmock.New()
	if err != nil {
		return
	}

	read, readMock, err := sqlmock.New()
	if err != nil {
		return
	}

	readMock.ExpectPrepare(accQueryFindByID)
	writeMock.ExpectPrepare(accQueryInsert)
	writeMock.ExpectPrepare(accQueryUpdate)
	writeMock.ExpectPrepare(accQueryDelete)

	resp, err = pg.NewAccountRepository(context.Background(), &pg.Config{
		PrimaryDB: write,
		ReplicaDB: read,
	})
	if err != nil {
		return
	}

	return resp, writeMock, readMock, nil
}

func randomAccStatus() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return accStatus[rand.Intn(len(accStatus))] // Pick a random status
}

func createRandomAccount() *db.Account {
	return &db.Account{
		ID:        faker.UUIDHyphenated(),
		Username:  faker.Username(),
		Email:     aws.String(faker.Email()),
		Phone:     aws.String(faker.Phonenumber()),
		Status:    randomAccStatus(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func argAccToArrDV(acc *db.Account) []driver.Value {
	return []driver.Value{
		acc.ID,
		acc.Username,
		acc.Email,
		acc.Phone,
		acc.Status,
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
	}
}

func accToArrDV(acc *db.Account) []driver.Value {
	return []driver.Value{
		acc.ID,
		acc.Username,
		acc.Email,
		acc.Phone,
		acc.Status,
		acc.CreatedAt.UTC(),
		acc.UpdatedAt.UTC(),
	}
}

func createRandomAccounts(i int) []*db.Account {
	var accounts []*db.Account
	for j := 0; j < i; j++ {
		accounts = append(accounts, createRandomAccount())
	}
	return accounts
}

var accList = createRandomAccounts(10)
