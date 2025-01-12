package pg_test

import (
	"log"
	"testing"
)

func TestAccountRepository_CreateAccount(t *testing.T) {
	// Unit Test
	UT_AccountRepository_CreateAccount(t)

	if testing.Short() {
		t.Log("skipping integration test in short mode.")
		return
	}

	log.Println("running integration test...")
}

func TestAccountRepository_FindAccountByID(t *testing.T) {
	// Unit Test
	UT_AccountRepository_FindAccountByID(t)

	if testing.Short() {
		t.Log("skipping integration test in short mode.")
		return
	}

	log.Println("running integration test...")
}

func TestAccountRepository_UpdateAccount(t *testing.T) {
	// Unit Test
	UT_AccountRepository_UpdateAccount(t)

	if testing.Short() {
		t.Log("skipping integration test in short mode.")
		return
	}

	log.Println("running integration test...")
}
