package pg

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nyaruka/phonenumbers"
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
	return newAccountRepository(ctx, config)
}

func newAccountRepository(ctx context.Context, config *Config) (*AccountRepository, error) {
	write, read, err := initConnection(config)
	if err != nil {
		return nil, err
	}

	stmtFindAccountByID, err := read.PrepareNamedContext(ctx, "SELECT id, username, email, phone, status, created_at, updated_at FROM accounts WHERE id = :id")
	if err != nil {
		return nil, err
	}

	stmtCreateAccount, err := write.PrepareNamedContext(ctx, "INSERT INTO accounts (id, username, email, phone, status, created_at, updated_at) VALUES (:id, :username, :email, :phone, :status, :created_at, :updated_at) RETURNING id, username, email, phone, status, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtUpdateAccount, err := write.PrepareNamedContext(ctx, "UPDATE accounts SET username = :username, email = :email, phone = :phone, status = :status, updated_at = :updated_at WHERE id = :id RETURNING id, username, email, phone, status, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtDeleteAccountByID, err := write.PrepareNamedContext(ctx, "DELETE FROM accounts WHERE id = :id")
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

// TODO: Not yet
func (this *AccountRepository) GetAccounts(ctx context.Context, param *db.ParamGetAccounts) (*db.Accounts, error) {
	return nil, errors.New("unimplemented")
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

	dbResp := accountEntity{}

	err := row.StructScan(&dbResp)
	if err != nil {
		return nil, err
	}

	resp := mapAccountEntityToAccount(&dbResp)

	return resp, nil
}

func validateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func validatePhone(phone string) error {
	parsedNumber, err := phonenumbers.Parse(phone, "")
	if err != nil {
		return fmt.Errorf("invalid phone number format: %v", err)
	}

	// Check if the number is valid
	if !phonenumbers.IsValidNumber(parsedNumber) {
		return fmt.Errorf("phone number is not valid")
	}

	return nil
}

func (this *AccountRepository) validateCreateAccount(ctx context.Context, param *db.ParamCreateAccount) error {
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

	if param.Phone != nil && (*param.Phone == "" || validatePhone(*param.Phone) != nil) {
		errs = append(errs, ErrInvalidPhone.Error())
	}

	if len(errs) > 0 {
		return ErrValidationFailed(errs)
	}

	return nil
}

func (this *AccountRepository) CreateAccount(ctx context.Context, param *db.ParamCreateAccount) (*db.Account, error) {
	if err := this.validateCreateAccount(ctx, param); err != nil {
		return nil, err
	}

	dbParam := mapParamCreateAccount(param)

	row := this.stmtCreateAccount.QueryRowxContext(ctx, dbParam)
	if err := row.Err(); err != nil {
		return nil, err
	}

	dbResp := accountEntity{}

	if err := row.StructScan(&dbResp); err != nil {
		return nil, err
	}

	resp := mapAccountEntityToAccount(&dbResp)

	return resp, nil
}

func (this *AccountRepository) validateUpdateAccount(ctx context.Context, param *db.ParamUpdateAccount) error {
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

func (this *AccountRepository) UpdateAccount(ctx context.Context, param *db.ParamUpdateAccount) (*db.Account, error) {
	if err := this.validateUpdateAccount(ctx, param); err != nil {
		return nil, err
	}

	dbParam := mapParamUpdateAccount(param)

	row := this.stmtUpdateAccount.QueryRowxContext(ctx, dbParam)

	if err := row.Err(); err != nil {
		return nil, err
	}

	dbResp := accountEntity{}

	if err := row.StructScan(&dbResp); err != nil {
		return nil, err
	}

	resp := mapAccountEntityToAccount(&dbResp)

	return resp, nil
}

func (this *AccountRepository) validatePatchAccount(param *db.ParamPatchAccount) error {
	if param.ID == "" {
		return ErrIdRequired
	}

	if param.Username == nil && param.Email == nil && param.Phone == nil && param.Status == nil {
		return ErrNoFieldToUpdate
	}

	if param.Email != nil && (*param.Email == "" || !validateEmail(*param.Email)) {
		return ErrInvalidEmail
	}

	if param.Phone != nil && (*param.Phone == "" || validatePhone(*param.Phone) != nil) {
		return ErrInvalidPhone
	}

	return nil
}

func (this *AccountRepository) PatchAccount(ctx context.Context, param *db.ParamPatchAccount) error {
	if err := this.validatePatchAccount(param); err != nil {
		return err
	}

	updateQuery := "UPDATE accounts SET "
	// build query based on the parameter
	queryArr := make([]string, 0)
	queryArgs := make(map[string]interface{})
	if param.Username != nil {
		if *param.Username == "" {
			queryArr = append(queryArr, "username = NULL")
		} else {
			queryArr = append(queryArr, "username = :username")
			queryArgs["username"] = *param.Username
		}
	}

	if param.Email != nil {
		if *param.Email == "" {
			queryArr = append(queryArr, "email = NULL")
		} else {
			queryArr = append(queryArr, "email = :email")
			queryArgs["email"] = *param.Email
		}
	}

	if param.Phone != nil {
		if *param.Phone == "" {
			queryArr = append(queryArr, "phone = NULL")
		} else {
			queryArr = append(queryArr, "phone = :phone")
			queryArgs["phone"] = *param.Phone
		}
	}

	if param.Status != nil {
		if *param.Status == "" {
			queryArr = append(queryArr, "status = NULL")
		} else {
			queryArr = append(queryArr, "status = :status")
			queryArgs["status"] = *param.Status
		}
	}

	if len(queryArr) == 0 {
		return ErrNoFieldToUpdate
	}

	queryArr = append(queryArr, "updated_at = :updated_at")
	queryArgs["updated_at"] = time.Now()

	updateQuery += strings.Join(queryArr, ", ") + " WHERE id = :id"

	if _, err := this.prim.NamedExecContext(ctx, updateQuery, queryArgs); err != nil {
		return err
	}

	return nil
}

func (this *AccountRepository) DeleteAccountByID(ctx context.Context, id string) error {
	if _, err := this.stmtDeleteAccountByID.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}

func (this *AccountRepository) beginTx(ctx context.Context, tx *sqlx.Tx) (db.IAccountTx, error) {
	if tx == nil {
		var err error
		tx, err = this.prim.BeginTxx(ctx, nil)
		if err != nil {
			return nil, err
		}
	}

	stmtFindAccountByID, err := tx.PrepareNamedContext(ctx, "SELECT id, username, email, phone, status, created_at, updated_at FROM accounts WHERE id = :id")
	if err != nil {
		return nil, err
	}

	stmtCreateAccount, err := tx.PrepareNamedContext(ctx, "INSERT INTO accounts (id, username, email, phone, status, created_at, updated_at) VALUES (:id, :username, :email, :phone, :status, :created_at, :updated_at) RETURNING id, username, email, phone, status, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtUpdateAccount, err := tx.PrepareNamedContext(ctx, "UPDATE accounts SET username = :username, email = :email, phone = :phone, status = :status WHERE id = :id RETURNING id, username, email, phone, status, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtDeleteAccountByID, err := tx.PrepareNamedContext(ctx, "DELETE FROM accounts WHERE id = :id")
	if err != nil {
		return nil, err
	}

	return &AccountTx{
		tx: tx,

		stmtFindAccountByID:   stmtFindAccountByID,
		stmtCreateAccount:     stmtCreateAccount,
		stmtUpdateAccount:     stmtUpdateAccount,
		stmtDeleteAccountByID: stmtDeleteAccountByID,
	}, nil
}

func (this *AccountRepository) BeginTx(ctx context.Context) (db.IAccountTx, error) {
	return this.beginTx(ctx, nil)
}

type AccountTx struct {
	tx *sqlx.Tx

	stmtFindAccountByID   *sqlx.NamedStmt
	stmtCreateAccount     *sqlx.NamedStmt
	stmtUpdateAccount     *sqlx.NamedStmt
	stmtDeleteAccountByID *sqlx.NamedStmt
}

func (this *AccountTx) Commit(ctx context.Context) error {
	return this.tx.Commit()
}

func (this *AccountTx) Rollback(ctx context.Context) error {
	return this.tx.Rollback()
}

func (this *AccountTx) CreateAccount(ctx context.Context, param *db.ParamCreateAccount) (*db.Account, error) {
	dbParam := mapParamCreateAccount(param)

	row := this.stmtCreateAccount.QueryRowxContext(ctx, dbParam)
	if err := row.Err(); err != nil {
		return nil, err
	}

	dbResp := accountEntity{}

	if err := row.StructScan(&dbResp); err != nil {
		return nil, err
	}

	resp := mapAccountEntityToAccount(&dbResp)

	return resp, nil
}

func (this *AccountTx) UpdateAccount(ctx context.Context, param *db.ParamUpdateAccount) (*db.Account, error) {
	dbParam := mapParamUpdateAccount(param)

	row := this.stmtUpdateAccount.QueryRowxContext(ctx, dbParam)

	if err := row.Err(); err != nil {
		return nil, err
	}

	dbResp := accountEntity{}

	if err := row.StructScan(&dbResp); err != nil {
		return nil, err
	}

	resp := mapAccountEntityToAccount(&dbResp)

	return resp, nil
}

func (this *AccountTx) PatchAccount(ctx context.Context, param *db.ParamPatchAccount) error {
	updateQuery := "UPDATE accounts SET "
	// build query based on the parameter
	queryArr := make([]string, 0)
	queryArgs := make(map[string]interface{})
	if param.Username != nil {
		if *param.Username == "" {
			queryArr = append(queryArr, "username = NULL")
		} else {
			queryArr = append(queryArr, "username = :username")
			queryArgs["username"] = *param.Username
		}
	}

	if param.Email != nil {
		if *param.Email == "" {
			queryArr = append(queryArr, "email = NULL")
		} else {
			queryArr = append(queryArr, "email = :email")
			queryArgs["email"] = *param.Email
		}
	}

	if param.Phone != nil {
		if *param.Phone == "" {
			queryArr = append(queryArr, "phone = NULL")
		} else {
			queryArr = append(queryArr, "phone = :phone")
			queryArgs["phone"] = *param.Phone
		}
	}

	if param.Status != nil {
		if *param.Status == "" {
			queryArr = append(queryArr, "status = NULL")
		} else {
			queryArr = append(queryArr, "status = :status")
			queryArgs["status"] = *param.Status
		}
	}

	if len(queryArr) == 0 {
		return ErrNoFieldToUpdate
	}

	queryArr = append(queryArr, "updated_at = :updated_at")
	queryArgs["updated_at"] = time.Now()

	updateQuery += strings.Join(queryArr, ", ") + " WHERE id = :id"

	if _, err := this.tx.NamedExecContext(ctx, updateQuery, queryArgs); err != nil {
		return err
	}

	return nil
}

func (this *AccountTx) DeleteAccountByID(ctx context.Context, id string) error {
	if _, err := this.stmtDeleteAccountByID.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
