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
)

var (
	accColumns = []string{"id", "username", "email", "phone", "status", "created_at", "updated_at"}
	accStatus  = []string{"active", "inactive"}
)

// newAccountRepositoryWithMock creates a new AccountRepository with mock database connection (write and read)
func newAccountRepositoryWithMock() (resp *pg.AccountRepository, writeMock sqlmock.Sqlmock, readMock sqlmock.Sqlmock, err error) {
	write, writeMock, err := sqlmock.New()
	if err != nil {
		return
	}

	read, readMock, err := sqlmock.New()
	if err != nil {
		return
	}

	readMock.ExpectPrepare(accQueryFindByID)

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

func randomAccount() *db.Account {
	return &db.Account{
		ID:        faker.UUIDHyphenated(),
		Username:  faker.Username(),
		Email:     aws.String(faker.Email()),
		Phone:     aws.String(faker.Phonenumber()),
		Status:    randomAccStatus(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func accToArrDV(acc *db.Account) []driver.Value {
	return []driver.Value{
		acc.ID,
		acc.Username,
		acc.Email,
		acc.Phone,
		acc.Status,
		acc.CreatedAt,
		acc.UpdatedAt,
	}
}

func generateListOfRandomAccounts(i int) []*db.Account {
	var accounts []*db.Account
	for j := 0; j < i; j++ {
		accounts = append(accounts, randomAccount())
	}
	return accounts
}

var accList = generateListOfRandomAccounts(10)

func randomAccFromList() *db.Account {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return accList[rand.Intn(len(accList))] // Pick a random account from the list
}
