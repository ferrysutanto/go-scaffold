package mock

import (
	"context"
	"reflect"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

func NewAccountRepositoryMock() *AccountRepositoryMock {
	return &AccountRepositoryMock{
		getAccountsExps:       make([]*getAccountsExp, 0),
		findAccountByIdExps:   make([]*getAccountByIdExp, 0),
		createAccountExps:     make([]*createAccountExp, 0),
		updateAccountExps:     make([]*updateAccountExp, 0),
		patchAccountExps:      make([]*patchAccountExp, 0),
		deleteAccountByIdExps: make([]*deleteAccountByIdExp, 0),
	}
}

type AccountRepositoryMock struct {
	getAccountsExps       []*getAccountsExp
	findAccountByIdExps   []*getAccountByIdExp
	createAccountExps     []*createAccountExp
	updateAccountExps     []*updateAccountExp
	patchAccountExps      []*patchAccountExp
	deleteAccountByIdExps []*deleteAccountByIdExp
	beginTxExps           []*beginTxExp
}

type getAccountsExp struct {
	param *db.ParamGetAccounts
	resp  *db.Accounts
	err   error
}

func (this *AccountRepositoryMock) ExpectGetAccounts(ctx context.Context, param *db.ParamGetAccounts, resp *db.Accounts, err error) {
	this.getAccountsExps = append(this.getAccountsExps, &getAccountsExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}
func (this *AccountRepositoryMock) GetAccounts(ctx context.Context, param *db.ParamGetAccounts) (*db.Accounts, error) {
	if len(this.getAccountsExps) == 0 {
		return nil, nil
	}

	for i, exp := range this.getAccountsExps {
		if !reflect.DeepEqual(exp.param, param) {
			continue
		}

		// remove the first element
		this.getAccountsExps = append(this.getAccountsExps[:i], this.getAccountsExps[i+1:]...)

		return exp.resp, exp.err
	}

	return nil, ErrExpectationNotFound
}

type getAccountByIdExp struct {
	id   string
	resp *db.Account
	err  error
}

func (this *AccountRepositoryMock) ExpectFindAccountByID(ctx context.Context, id string, resp *db.Account, err error) {
	this.findAccountByIdExps = append(this.findAccountByIdExps, &getAccountByIdExp{
		id:   id,
		resp: resp,
		err:  err,
	})
}

func (this *AccountRepositoryMock) FindAccountByID(ctx context.Context, id string) (*db.Account, error) {
	if len(this.findAccountByIdExps) == 0 {
		return nil, nil
	}

	for i, exp := range this.findAccountByIdExps {
		if exp.id != id {
			continue
		}

		// remove the first element
		this.findAccountByIdExps = append(this.findAccountByIdExps[:i], this.findAccountByIdExps[i+1:]...)

		return exp.resp, exp.err
	}

	return nil, ErrExpectationNotFound
}

func (this *AccountRepositoryMock) ExpectCreateAccount(ctx context.Context, param *db.ParamCreateAccount, resp *db.Account, err error) {
	this.createAccountExps = append(this.createAccountExps, &createAccountExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}

func (this *AccountRepositoryMock) CreateAccount(ctx context.Context, param *db.ParamCreateAccount) (*db.Account, error) {
	if len(this.createAccountExps) == 0 {
		return nil, nil
	}

	for i, exp := range this.createAccountExps {
		// check if the input is the same
		if !reflect.DeepEqual(exp.param, param) {
			continue
		}

		// remove the first element
		this.createAccountExps = append(this.createAccountExps[:i], this.createAccountExps[i+1:]...)

		return exp.resp, exp.err
	}

	return nil, ErrExpectationNotFound
}

func (this *AccountRepositoryMock) ExpectUpdateAccount(ctx context.Context, param *db.ParamUpdateAccount, resp *db.Account, err error) {
	this.updateAccountExps = append(this.updateAccountExps, &updateAccountExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}

func (this *AccountRepositoryMock) UpdateAccount(ctx context.Context, param *db.ParamUpdateAccount) (*db.Account, error) {
	if len(this.updateAccountExps) == 0 {
		return nil, nil
	}

	for i, exp := range this.updateAccountExps {
		// check if the input is the same
		if !reflect.DeepEqual(exp.param, param) {
			continue
		}

		// remove the first element
		this.updateAccountExps = append(this.updateAccountExps[:i], this.updateAccountExps[i+1:]...)

		return exp.resp, exp.err
	}

	return nil, ErrExpectationNotFound
}

func (this *AccountRepositoryMock) ExpectPatchAccount(ctx context.Context, param *db.ParamPatchAccount, err error) {
	this.patchAccountExps = append(this.patchAccountExps, &patchAccountExp{
		param: param,
		err:   err,
	})
}

func (this *AccountRepositoryMock) PatchAccount(ctx context.Context, param *db.ParamPatchAccount) error {
	if len(this.patchAccountExps) == 0 {
		return nil
	}

	for i, exp := range this.patchAccountExps {
		// check if the input is the same
		if !reflect.DeepEqual(exp.param, param) {
			continue
		}

		// remove the first element
		this.patchAccountExps = append(this.patchAccountExps[:i], this.patchAccountExps[i+1:]...)

		return exp.err
	}

	return ErrExpectationNotFound
}

func (this *AccountRepositoryMock) ExpectDeleteAccountByID(ctx context.Context, id string, err error) {
	this.deleteAccountByIdExps = append(this.deleteAccountByIdExps, &deleteAccountByIdExp{
		id:  id,
		err: err,
	})
}

func (this *AccountRepositoryMock) DeleteAccountByID(ctx context.Context, id string) error {
	if len(this.deleteAccountByIdExps) == 0 {
		return nil
	}

	for i, exp := range this.deleteAccountByIdExps {
		// check if the input is the same
		if exp.id != id {
			continue
		}

		// remove the first element
		this.deleteAccountByIdExps = append(this.deleteAccountByIdExps[:i], this.deleteAccountByIdExps[i+1:]...)

		return exp.err
	}

	return ErrExpectationNotFound
}

type beginTxExp struct {
	resp db.IAccountTx
	err  error
}

func (this *AccountRepositoryMock) ExpectBeginTx(ctx context.Context, resp db.IAccountTx, err error) {
	this.beginTxExps = append(this.beginTxExps, &beginTxExp{
		resp: resp,
		err:  err,
	})
}

func (this *AccountRepositoryMock) BeginTx(ctx context.Context) (db.IAccountTx, error) {
	if len(this.beginTxExps) == 0 {
		return nil, ErrFulfilled
	}

	exp := this.beginTxExps[0]
	this.beginTxExps = this.beginTxExps[1:]

	return exp.resp, exp.err
}
