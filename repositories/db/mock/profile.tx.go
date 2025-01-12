package mock

import (
	"context"
	"reflect"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

func NewProfileTxMock() *ProfileTxMock {
	return &ProfileTxMock{
		commitExps:            make([]*commitExp, 0),
		rollbackExps:          make([]*rollbackExp, 0),
		createProfileExps:     make([]*createProfileExp, 0),
		updateProfileExps:     make([]*updateProfileExp, 0),
		patchProfileExps:      make([]*patchProfileExp, 0),
		deleteProfileByIdExps: make([]*deleteProfileByIdExp, 0),
	}
}

type ProfileTxMock struct {
	commitExps            []*commitExp
	rollbackExps          []*rollbackExp
	createProfileExps     []*createProfileExp
	updateProfileExps     []*updateProfileExp
	patchProfileExps      []*patchProfileExp
	deleteProfileByIdExps []*deleteProfileByIdExp
}

func (this *ProfileTxMock) ExpectCommit(ctx context.Context, expResp error) {
	this.commitExps = append(this.commitExps, &commitExp{
		resp: expResp,
	})
}

func (this *ProfileTxMock) Commit(ctx context.Context) error {
	if len(this.commitExps) == 0 {
		return ErrFulfilled
	}

	exp := this.commitExps[0]

	this.commitExps = this.commitExps[1:]

	return exp.resp
}

func (this *ProfileTxMock) ExpectRollback(ctx context.Context, expResp error) {
	this.rollbackExps = append(this.rollbackExps, &rollbackExp{
		resp: expResp,
	})
}

func (this *ProfileTxMock) Rollback(ctx context.Context) error {
	if len(this.rollbackExps) == 0 {
		return ErrFulfilled
	}

	exp := this.rollbackExps[0]

	this.rollbackExps = this.rollbackExps[1:]

	return exp.resp
}

type createProfileExp struct {
	param *db.ParamCreateProfile
	resp  *db.Profile
	err   error
}

func (this *ProfileTxMock) ExpectCreateProfile(ctx context.Context, param *db.ParamCreateProfile, resp *db.Profile, err error) {
	this.createProfileExps = append(this.createProfileExps, &createProfileExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}

func (this *ProfileTxMock) CreateProfile(ctx context.Context, param *db.ParamCreateProfile) (*db.Profile, error) {
	if len(this.createProfileExps) == 0 {
		return nil, ErrFulfilled
	}

	for i, exp := range this.createProfileExps {
		if !reflect.DeepEqual(exp.param, param) {
			continue
		}

		this.createProfileExps = append(this.createProfileExps[:i], this.createProfileExps[i+1:]...)

		return exp.resp, exp.err
	}

	return nil, ErrExpectationNotFound
}

type updateProfileExp struct {
	param *db.ParamUpdateProfile
	resp  *db.Profile
	err   error
}

func (this *ProfileTxMock) ExpectUpdateProfile(ctx context.Context, param *db.ParamUpdateProfile, resp *db.Profile, err error) {
	this.updateProfileExps = append(this.updateProfileExps, &updateProfileExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}

func (this *ProfileTxMock) UpdateProfile(ctx context.Context, param *db.ParamUpdateProfile) (*db.Profile, error) {
	if len(this.updateProfileExps) == 0 {
		return nil, ErrFulfilled
	}

	for i, exp := range this.updateProfileExps {
		if !reflect.DeepEqual(exp.param, param) {
			continue
		}

		this.updateProfileExps = append(this.updateProfileExps[:i], this.updateProfileExps[i+1:]...)

		return exp.resp, exp.err
	}

	return nil, ErrExpectationNotFound
}

type patchProfileExp struct {
	param *db.ParamPatchProfile
	err   error
}

func (this *ProfileTxMock) ExpectPatchProfile(ctx context.Context, param *db.ParamPatchProfile, err error) {
	this.patchProfileExps = append(this.patchProfileExps, &patchProfileExp{
		param: param,
		err:   err,
	})
}

func (this *ProfileTxMock) PatchProfile(ctx context.Context, param *db.ParamPatchProfile) error {
	if len(this.patchProfileExps) == 0 {
		return ErrFulfilled
	}

	for i, exp := range this.patchProfileExps {
		if !reflect.DeepEqual(exp.param, param) {
			continue
		}

		this.patchProfileExps = append(this.patchProfileExps[:i], this.patchProfileExps[i+1:]...)

		return exp.err
	}

	return ErrExpectationNotFound
}

type deleteProfileByIdExp struct {
	id  string
	err error
}

func (this *ProfileTxMock) ExpectDeleteProfileByID(ctx context.Context, id string, err error) {
	this.deleteProfileByIdExps = append(this.deleteProfileByIdExps, &deleteProfileByIdExp{
		id:  id,
		err: err,
	})
}

func (this *ProfileTxMock) DeleteProfileByID(ctx context.Context, id string) error {
	if len(this.deleteProfileByIdExps) == 0 {
		return ErrFulfilled
	}

	for i, exp := range this.deleteProfileByIdExps {
		if exp.id != id {
			continue
		}

		this.deleteProfileByIdExps = append(this.deleteProfileByIdExps[:i], this.deleteProfileByIdExps[i+1:]...)

		return exp.err
	}

	return ErrExpectationNotFound
}
