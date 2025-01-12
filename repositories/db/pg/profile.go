package pg

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProfileRepository struct {
	prim *sqlx.DB
	repl *sqlx.DB

	// Only static queries are prepared, dynamic queries such as GetProfiles and PatchProfile are not prepared
	stmtFindProfileByID   *sqlx.NamedStmt
	stmtCreateProfile     *sqlx.NamedStmt
	stmtUpdateProfile     *sqlx.NamedStmt
	stmtDeleteProfileByID *sqlx.NamedStmt
}

func NewProfileRepository(ctx context.Context, config *Config) (db.IProfileRepository, error) {
	return newProfileRepository(ctx, config)
}

func newProfileRepository(ctx context.Context, config *Config) (*ProfileRepository, error) {
	write, read, err := initConnection(config)
	if err != nil {
		return nil, err
	}

	stmtFindProfileByID, err := read.PrepareNamedContext(ctx, "SELECT id, account_id, first_name, last_name, birthdate, sex, picture_url, created_at, updated_at FROM profiles WHERE id = :id")
	if err != nil {
		return nil, err
	}

	stmtCreateProfile, err := write.PrepareNamedContext(ctx, "INSERT INTO profiles (id, account_id, first_name, last_name, birthdate, sex, picture_url, created_at, updated_at) VALUES (:id, :account_id, :first_name, :last_name, :birthdate, :sex, :picture_url, :created_at, :updated_at) RETURNING id, account_id, first_name, last_name, birthdate, sex, picture_url, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtUpdateProfile, err := write.PrepareNamedContext(ctx, "UPDATE profiles SET account_id = :account_id, first_name = :first_name, last_name = :last_name, birthdate = :birthdate, sex = :sex, picture_url = :picture_url, updated_at = :updated_at WHERE id = :id RETURNING id, account_id, first_name, last_name, birthdate, sex, picture_url, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtDeleteProfileByID, err := write.PrepareNamedContext(ctx, "DELETE FROM profiles WHERE id = :id")
	if err != nil {
		return nil, err
	}

	return &ProfileRepository{
		prim: write,
		repl: read,

		stmtFindProfileByID:   stmtFindProfileByID,
		stmtCreateProfile:     stmtCreateProfile,
		stmtUpdateProfile:     stmtUpdateProfile,
		stmtDeleteProfileByID: stmtDeleteProfileByID,
	}, nil
}

// SIGNATURE SECTION

// TODO: Not yet
func (this *ProfileRepository) GetProfiles(ctx context.Context, param *db.ParamGetProfiles) (*db.Profiles, error) {
	return nil, errors.New("unimplemented")
}

func (this *ProfileRepository) validateFindProfileByID(ctx context.Context, id string) error {
	if id == "" {
		return ErrIdRequired
	}

	//  if not UUID
	if _, err := uuid.Parse(id); err != nil {
		return ErrInvalidUUID
	}

	return nil
}
func (this *ProfileRepository) FindProfileByID(ctx context.Context, id string) (*db.Profile, error) {
	if err := this.validateFindProfileByID(ctx, id); err != nil {
		return nil, err
	}

	row := this.stmtFindProfileByID.QueryRowxContext(ctx, map[string]interface{}{"id": id})
	if row.Err() != nil {
		return nil, row.Err()
	}

	dbResp := profileEntity{}

	err := row.StructScan(&dbResp)
	if err != nil {
		return nil, err
	}

	resp := mapProfileEntityToProfile(&dbResp)

	return resp, nil
}

func (this *ProfileRepository) validateCreateProfile(ctx context.Context, param *db.ParamCreateProfile) error {
	errs := make([]string, 0)

	if len(errs) > 0 {
		return ErrValidationFailed(errs)
	}

	return nil
}

func (this *ProfileRepository) CreateProfile(ctx context.Context, param *db.ParamCreateProfile) (*db.Profile, error) {
	if err := this.validateCreateProfile(ctx, param); err != nil {
		return nil, err
	}

	dbParam := mapParamCreateProfile(param)

	row := this.stmtCreateProfile.QueryRowxContext(ctx, dbParam)
	if err := row.Err(); err != nil {
		return nil, err
	}

	dbResp := profileEntity{}

	if err := row.StructScan(&dbResp); err != nil {
		return nil, err
	}

	resp := mapProfileEntityToProfile(&dbResp)

	return resp, nil
}

func (this *ProfileRepository) validateUpdateProfile(ctx context.Context, param *db.ParamUpdateProfile) error {
	errs := make([]string, 0)

	if len(errs) > 0 {
		return ErrValidationFailed(errs)
	}

	return nil
}

func (this *ProfileRepository) UpdateProfile(ctx context.Context, param *db.ParamUpdateProfile) (*db.Profile, error) {
	if err := this.validateUpdateProfile(ctx, param); err != nil {
		return nil, err
	}

	dbParam := mapParamUpdateProfile(param)

	row := this.stmtUpdateProfile.QueryRowxContext(ctx, dbParam)

	if err := row.Err(); err != nil {
		return nil, err
	}

	dbResp := profileEntity{}

	if err := row.StructScan(&dbResp); err != nil {
		return nil, err
	}

	resp := mapProfileEntityToProfile(&dbResp)

	return resp, nil
}

func (this *ProfileRepository) validatePatchProfile(param *db.ParamPatchProfile) error {
	if param.ID == "" {
		return ErrIdRequired
	}

	return nil
}

func (this *ProfileRepository) PatchProfile(ctx context.Context, param *db.ParamPatchProfile) error {
	if err := this.validatePatchProfile(param); err != nil {
		return err
	}

	updateQuery := "UPDATE Profiles SET "
	// build query based on the parameter
	queryArr := make([]string, 0)
	queryArgs := make(map[string]interface{})

	if len(queryArr) == 0 {
		return ErrNoFieldToUpdate
	}

	if param.FirstName != nil {
		queryArr = append(queryArr, "first_name = :first_name")
		queryArgs["first_name"] = *param.FirstName
	}

	if param.LastName != nil {
		queryArr = append(queryArr, "last_name = :last_name")
		queryArgs["last_name"] = *param.LastName
	}

	if param.Birthdate != nil {
		queryArr = append(queryArr, "birthdate = :birthdate")
		queryArgs["birthdate"] = *param.Birthdate
	}

	if param.Sex != nil {
		queryArr = append(queryArr, "sex = :sex")
		queryArgs["sex"] = *param.Sex
	}

	if param.PictureURL != nil {
		queryArr = append(queryArr, "picture_url = :picture_url")
		queryArgs["picture_url"] = *param.PictureURL
	}

	queryArr = append(queryArr, "updated_at = :updated_at")
	queryArgs["updated_at"] = time.Now()

	updateQuery += strings.Join(queryArr, ", ") + " WHERE id = :id"

	if _, err := this.prim.NamedExecContext(ctx, updateQuery, queryArgs); err != nil {
		return err
	}

	return nil
}

func (this *ProfileRepository) DeleteProfileByID(ctx context.Context, id string) error {
	if _, err := this.stmtDeleteProfileByID.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}

func (this *ProfileRepository) beginTx(ctx context.Context, tx *sqlx.Tx) (db.IProfileTx, error) {
	if tx == nil {
		var err error
		tx, err = this.prim.BeginTxx(ctx, nil)
		if err != nil {
			return nil, err
		}
	}

	stmtFindProfileByID, err := tx.PrepareNamedContext(ctx, "SELECT id, username, email, phone, status, created_at, updated_at FROM Profiles WHERE id = :id")
	if err != nil {
		return nil, err
	}

	stmtCreateProfile, err := tx.PrepareNamedContext(ctx, "INSERT INTO Profiles (id, username, email, phone, status, created_at, updated_at) VALUES (:id, :username, :email, :phone, :status, :created_at, :updated_at) RETURNING id, username, email, phone, status, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtUpdateProfile, err := tx.PrepareNamedContext(ctx, "UPDATE Profiles SET username = :username, email = :email, phone = :phone, status = :status WHERE id = :id RETURNING id, username, email, phone, status, created_at, updated_at")
	if err != nil {
		return nil, err
	}

	stmtDeleteProfileByID, err := tx.PrepareNamedContext(ctx, "DELETE FROM Profiles WHERE id = :id")
	if err != nil {
		return nil, err
	}

	return &ProfileTx{
		tx: tx,

		stmtFindProfileByID:   stmtFindProfileByID,
		stmtCreateProfile:     stmtCreateProfile,
		stmtUpdateProfile:     stmtUpdateProfile,
		stmtDeleteProfileByID: stmtDeleteProfileByID,
	}, nil
}

func (this *ProfileRepository) BeginTx(ctx context.Context) (db.IProfileTx, error) {
	return this.beginTx(ctx, nil)
}

type ProfileTx struct {
	tx *sqlx.Tx

	stmtFindProfileByID   *sqlx.NamedStmt
	stmtCreateProfile     *sqlx.NamedStmt
	stmtUpdateProfile     *sqlx.NamedStmt
	stmtDeleteProfileByID *sqlx.NamedStmt
}

func (this *ProfileTx) Commit(ctx context.Context) error {
	return this.tx.Commit()
}

func (this *ProfileTx) Rollback(ctx context.Context) error {
	return this.tx.Rollback()
}

func (this *ProfileTx) CreateProfile(ctx context.Context, param *db.ParamCreateProfile) (*db.Profile, error) {
	dbParam := mapParamCreateProfile(param)

	row := this.stmtCreateProfile.QueryRowxContext(ctx, dbParam)
	if err := row.Err(); err != nil {
		return nil, err
	}

	dbResp := profileEntity{}

	if err := row.StructScan(&dbResp); err != nil {
		return nil, err
	}

	resp := mapProfileEntityToProfile(&dbResp)

	return resp, nil
}

func (this *ProfileTx) UpdateProfile(ctx context.Context, param *db.ParamUpdateProfile) (*db.Profile, error) {
	dbParam := mapParamUpdateProfile(param)

	row := this.stmtUpdateProfile.QueryRowxContext(ctx, dbParam)

	if err := row.Err(); err != nil {
		return nil, err
	}

	dbResp := profileEntity{}

	if err := row.StructScan(&dbResp); err != nil {
		return nil, err
	}

	resp := mapProfileEntityToProfile(&dbResp)

	return resp, nil
}

func (this *ProfileTx) validatePatchProfile(param *db.ParamPatchProfile) error {
	if param.ID == "" {
		return ErrIdRequired
	}

	return nil
}

func (this *ProfileTx) PatchProfile(ctx context.Context, param *db.ParamPatchProfile) error {
	if err := this.validatePatchProfile(param); err != nil {
		return err
	}

	updateQuery := "UPDATE Profiles SET "
	// build query based on the parameter
	queryArr := make([]string, 0)
	queryArgs := make(map[string]interface{})

	if len(queryArr) == 0 {
		return ErrNoFieldToUpdate
	}

	if param.FirstName != nil {
		queryArr = append(queryArr, "first_name = :first_name")
		queryArgs["first_name"] = *param.FirstName
	}

	if param.LastName != nil {
		queryArr = append(queryArr, "last_name = :last_name")
		queryArgs["last_name"] = *param.LastName
	}

	if param.Birthdate != nil {
		queryArr = append(queryArr, "birthdate = :birthdate")
		queryArgs["birthdate"] = *param.Birthdate
	}

	if param.Sex != nil {
		queryArr = append(queryArr, "sex = :sex")
		queryArgs["sex"] = *param.Sex
	}

	if param.PictureURL != nil {
		queryArr = append(queryArr, "picture_url = :picture_url")
		queryArgs["picture_url"] = *param.PictureURL
	}

	queryArr = append(queryArr, "updated_at = :updated_at")
	queryArgs["updated_at"] = time.Now()

	updateQuery += strings.Join(queryArr, ", ") + " WHERE id = :id"

	if _, err := this.tx.NamedExecContext(ctx, updateQuery, queryArgs); err != nil {
		return err
	}

	return nil
}

func (this *ProfileTx) DeleteProfileByID(ctx context.Context, id string) error {
	if _, err := this.stmtDeleteProfileByID.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
