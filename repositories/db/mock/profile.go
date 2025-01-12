package mock

import (
	"context"
	"reflect"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

func NewProfileRepositoryMock() *ProfileRepositoryMock {
	return &ProfileRepositoryMock{
		getProfilesExps:       make([]*getProfilesExp, 0),
		findProfileByIdExps:   make([]*getProfileByIdExp, 0),
		createProfileExps:     make([]*createProfileExp, 0),
		updateProfileExps:     make([]*updateProfileExp, 0),
		patchProfileExps:      make([]*patchProfileExp, 0),
		deleteProfileByIdExps: make([]*deleteProfileByIdExp, 0),
	}
}

type ProfileRepositoryMock struct {
	getProfilesExps       []*getProfilesExp
	findProfileByIdExps   []*getProfileByIdExp
	createProfileExps     []*createProfileExp
	updateProfileExps     []*updateProfileExp
	patchProfileExps      []*patchProfileExp
	deleteProfileByIdExps []*deleteProfileByIdExp
}

type getProfilesExp struct {
	param *db.ParamGetProfiles
	resp  *db.Profiles
	err   error
}

func (this *ProfileRepositoryMock) ExpectGetProfiles(ctx context.Context, param *db.ParamGetProfiles, resp *db.Profiles, err error) {
	this.getProfilesExps = append(this.getProfilesExps, &getProfilesExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}

func (this *ProfileRepositoryMock) GetProfiles(ctx context.Context, param *db.ParamGetProfiles) (*db.Profiles, error) {
	if len(this.getProfilesExps) == 0 {
		return nil, nil
	}

	for i, exp := range this.getProfilesExps {
		if !reflect.DeepEqual(exp.param, param) {
			continue
		}

		this.getProfilesExps = append(this.getProfilesExps[:i], this.getProfilesExps[i+1:]...)

		return exp.resp, exp.err
	}

	return nil, ErrExpectationNotFound
}

type getProfileByIdExp struct {
	id   string
	resp *db.Profile
	err  error
}

func (this *ProfileRepositoryMock) ExpectFindProfileByID(ctx context.Context, id string, resp *db.Profile, err error) {
	this.findProfileByIdExps = append(this.findProfileByIdExps, &getProfileByIdExp{
		id:   id,
		resp: resp,
		err:  err,
	})
}

func (this *ProfileRepositoryMock) FindProfileByID(ctx context.Context, id string) (*db.Profile, error) {
	if len(this.findProfileByIdExps) == 0 {
		return nil, nil
	}

	for i, exp := range this.findProfileByIdExps {
		if exp.id != id {
			continue
		}

		this.findProfileByIdExps = append(this.findProfileByIdExps[:i], this.findProfileByIdExps[i+1:]...)

		return exp.resp, exp.err
	}

	return nil, ErrExpectationNotFound
}

func (this *ProfileRepositoryMock) ExpectCreateProfile(ctx context.Context, param *db.ParamCreateProfile, resp *db.Profile, err error) {
	this.createProfileExps = append(this.createProfileExps, &createProfileExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}

func (this *ProfileRepositoryMock) CreateProfile(ctx context.Context, param *db.ParamCreateProfile) (*db.Profile, error) {
	if len(this.createProfileExps) == 0 {
		return nil, nil
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

func (this *ProfileRepositoryMock) ExpectUpdateProfile(ctx context.Context, param *db.ParamUpdateProfile, resp *db.Profile, err error) {
	this.updateProfileExps = append(this.updateProfileExps, &updateProfileExp{
		param: param,
		resp:  resp,
		err:   err,
	})
}

func (this *ProfileRepositoryMock) UpdateProfile(ctx context.Context, param *db.ParamUpdateProfile) (*db.Profile, error) {
	if len(this.updateProfileExps) == 0 {
		return nil, nil
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

func (this *ProfileRepositoryMock) ExpectPatchProfile(ctx context.Context, param *db.ParamPatchProfile, err error) {
	this.patchProfileExps = append(this.patchProfileExps, &patchProfileExp{
		param: param,
		err:   err,
	})
}

func (this *ProfileRepositoryMock) PatchProfile(ctx context.Context, param *db.ParamPatchProfile) error {
	if len(this.patchProfileExps) == 0 {
		return nil
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

func (this *ProfileRepositoryMock) ExpectDeleteProfileByID(ctx context.Context, id string, err error) {
	this.deleteProfileByIdExps = append(this.deleteProfileByIdExps, &deleteProfileByIdExp{
		id:  id,
		err: err,
	})
}

func (this *ProfileRepositoryMock) DeleteProfileByID(ctx context.Context, id string) error {
	if len(this.deleteProfileByIdExps) == 0 {
		return nil
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
