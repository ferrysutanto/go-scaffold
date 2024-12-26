package pg_test

import (
	"testing"
)

func TestAccountRepository_FindAccountByID(t *testing.T) {
	UT_AccountRepository_FindAccountByID(t)

	if testing.Short() {
		t.Log("skipping integration test in short mode.")
	}
}
