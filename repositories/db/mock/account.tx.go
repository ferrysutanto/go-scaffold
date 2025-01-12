package mock

import (
	"context"
	"reflect"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

func NewAccountTxMock() *AccountTxMock {
	return &AccountTxMock{
		commitExps:            make([]*commitExp, 0),
		rollbackExps:          make([]*rollbackExp, 0),
		createAccountExps:     make([]*createAccountExp, 0),
		updateAccountExps:     make([]*updateAccountExp, 0),
		patchAccountExps:      make([]*patchAccountExp, 0),
		deleteAccountByIdExps: make([]*deleteAccountByIdExp, 0),
	}
}

type AccountTxMock struct {
	commitExps            []*commitExp
	rollbackExps          []*rollbackExp
	createAccountExps     []*createAccountExp
	updateAccountExps     []*updateAccountExp
	patchAccountExps      []*patchAccountExp
	deleteAccountByIdExps []*deleteAccountByIdExp
}

type commitExp struct {
	resp error
}

func (this *AccountTxMock) ExpectCommit(ctx context.Context, expResp error) {
	this.commitExps = append(this.commitExps, &commitExp{
		resp: expResp,
	})
}

func (this *AccountTxMock) Commit(ctx context.Context) error {

	if len(this.commitExps) == 0 {
		return ErrFulfilled
	}

	exp := this.commitExps[0]

	// remove the first element
	this.commitExps = this.commitExps[1:]

	return exp.resp
}

type rollbackExp struct {
	resp error
}

func (this *AccountTxMock) ExpectRollback(ctx context.Context, expResp error) {
	this.rollbackExps = append(this.rollbackExps, &rollbackExp{
		resp: expResp,
	})
}
func (this *AccountTxMock) Rollback(ctx context.Context) error {
	if len(this.rollbackExps) == 0 {
		return ErrFulfilled
	}

	exp := this.rollbackExps[0]

	// remove the first element
	this.rollbackExps = this.rollbackExps[1:]

	return exp.resp
}

type createAccountExp struct {
	param *db.ParamCreateAccount
	resp  *db.Account
	err   error
}

func (this *AccountTxMock) ExpectCreateAccount(ctx context.Context, param *db.ParamCreateAccount, resp *db.Account, err error) {
	this.createAccountExps = append(this.createAccountExps, &createAccountExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}

func (this *AccountTxMock) CreateAccount(ctx context.Context, param *db.ParamCreateAccount) (*db.Account, error) {
	if len(this.createAccountExps) == 0 {
		return nil, ErrFulfilled
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

type updateAccountExp struct {
	param *db.ParamUpdateAccount
	resp  *db.Account
	err   error
}

func (this *AccountTxMock) ExpectUpdateAccount(ctx context.Context, param *db.ParamUpdateAccount, resp *db.Account, err error) {
	this.updateAccountExps = append(this.updateAccountExps, &updateAccountExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}
func (this *AccountTxMock) UpdateAccount(ctx context.Context, param *db.ParamUpdateAccount) (*db.Account, error) {
	if len(this.updateAccountExps) == 0 {
		return nil, ErrFulfilled
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

type patchAccountExp struct {
	param *db.ParamPatchAccount
	err   error
}

func (this *AccountTxMock) ExpectPatchAccount(ctx context.Context, param *db.ParamPatchAccount, err error) {
	this.patchAccountExps = append(this.patchAccountExps, &patchAccountExp{
		param: param,
		err:   err,
	})
}
func (this *AccountTxMock) PatchAccount(ctx context.Context, param *db.ParamPatchAccount) error {
	if len(this.patchAccountExps) == 0 {
		return ErrFulfilled
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

type deleteAccountByIdExp struct {
	id  string
	err error
}

func (this *AccountTxMock) ExpectDeleteAccountByID(ctx context.Context, id string, err error) {
	this.deleteAccountByIdExps = append(this.deleteAccountByIdExps, &deleteAccountByIdExp{
		id:  id,
		err: err,
	})
}
func (this *AccountTxMock) DeleteAccountByID(ctx context.Context, id string) error {
	if len(this.deleteAccountByIdExps) == 0 {
		return ErrFulfilled
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
