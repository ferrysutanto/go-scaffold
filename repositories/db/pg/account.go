package pg

import (
	"context"
	"regexp"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	prim *sqlx.DB
	repl *sqlx.DB

	// Only static queries are prepared, dynamic queries such as GetAccounts and PatchAccount are not prepared
	stmtFindAccountByID   *sqlx.NamedStmt
	stmtCreateAccount     *sqlx.NamedStmt
	stmtUpdateAccount     *sqlx.NamedStmt
	stmtDeleteAccountByID *sqlx.NamedStmt
}

func NewAccountRepository(ctx context.Context, config *Config) (db.IAccountRepository, error) {
	write, read, err := initConnection(config)
	if err != nil {
		return nil, err
	}

	stmtFindAccountByID, err := read.PrepareNamed("SELECT id, username, email, phone, status, created_at, updated_at FROM accounts WHERE id = :id")
	if err != nil {
		return nil, err
	}

	stmtCreateAccount, err := write.PrepareNamed("INSERT INTO accounts (id, username, email, phone, status, created_at, updated_at) VALUES (:id, :username, :email, :phone, :status, :created_at, :updated_at) RETURNING id, username, email, phone, status, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtUpdateAccount, err := write.PrepareNamed("UPDATE accounts SET username = :username, email = :email, phone = :phone, status = :status WHERE id = :id RETURNING id, username, email, phone, status, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtDeleteAccountByID, err := write.PrepareNamed("DELETE FROM accounts WHERE id = :id")
	if err != nil {
		return nil, err
	}

	return &AccountRepository{
		prim: write,
		repl: read,

		stmtFindAccountByID:   stmtFindAccountByID,
		stmtCreateAccount:     stmtCreateAccount,
		stmtUpdateAccount:     stmtUpdateAccount,
		stmtDeleteAccountByID: stmtDeleteAccountByID,
	}, nil
}

// SIGNATURE SECTION

func (repo *AccountRepository) GetAccounts(ctx context.Context, param *db.ParamGetAccounts) (*db.Accounts, error) {
	return nil, nil
}

func (this *AccountRepository) validateFindAccountByID(ctx context.Context, id string) error {
	if id == "" {
		return ErrIdRequired
	}

	//  if not UUID
	if _, err := uuid.Parse(id); err != nil {
		return ErrInvalidUUID
	}

	return nil
}
func (this *AccountRepository) FindAccountByID(ctx context.Context, id string) (*db.Account, error) {
	if err := this.validateFindAccountByID(ctx, id); err != nil {
		return nil, err
	}

	row := this.stmtFindAccountByID.QueryRowxContext(ctx, map[string]interface{}{"id": id})
	if row.Err() != nil {
		return nil, row.Err()
	}

	resp := db.Account{}

	err := row.StructScan(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func validateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func validateCreateAccount(ctx context.Context, param *db.ParamCreateAccount) error {
	errs := make([]string, 0)
	if param.Username == "" {
		errs = append(errs, ErrUsernameRequired.Error())
	}

	if param.Email == nil && param.Phone == nil {
		errs = append(errs, ErrEmailOrPhoneRequired.Error())
	}

	if param.Email != nil && (*param.Email == "" || !validateEmail(*param.Email)) {
		errs = append(errs, ErrInvalidEmail.Error())
	}

	// TODO: phone aren't handled yet

	if len(errs) > 0 {
		return ErrValidationFailed(errs)
	}

	return nil
}

func (this *AccountRepository) CreateAccount(ctx context.Context, param *db.ParamCreateAccount) (*db.Account, error) {
	if err := validateCreateAccount(ctx, param); err != nil {
		return nil, err
	}

	dbParam := mapParamCreateAccount(param)

	row := this.stmtCreateAccount.QueryRowxContext(ctx, dbParam)
	if row.Err() != nil {
		return nil, row.Err()
	}

	resp := db.Account{}

	if err := row.StructScan(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
func (repo *AccountRepository) UpdateAccount(ctx context.Context, param *db.ParamUpdateAccount) (*db.Account, error) {
	return nil, nil
}
func (repo *AccountRepository) PatchAccount(ctx context.Context, param *db.ParamPatchAccount) (*db.Account, error) {
	return nil, nil
}
func (repo *AccountRepository) DeleteAccountByID(ctx context.Context, id string) error {
	return nil
}

func (repo *AccountRepository) BeginTx(ctx context.Context) (db.IAccountTx, error) {
	return &AccountTx{}, nil
}

type AccountTx struct{}

func (tx *AccountTx) Commit(ctx context.Context) error { return nil }

func (tx *AccountTx) Rollback(ctx context.Context) error { return nil }

func (tx *AccountTx) CreateAccount(ctx context.Context, account *db.ParamCreateAccount) (*db.Account, error) {
	return nil, nil
}
func (tx *AccountTx) UpdateAccount(ctx context.Context, account *db.ParamUpdateAccount) (*db.Account, error) {
	return nil, nil
}
func (tx *AccountTx) PatchAccount(ctx context.Context, account *db.ParamPatchAccount) (*db.Account, error) {
	return nil, nil
}
func (tx *AccountTx) DeleteAccountByID(ctx context.Context, id string) error { return nil }
