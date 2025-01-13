package mock

import (
	"context"

	"github.com/ferrysutanto/go-scaffold/repositories/db"
)

type DB struct {
	pingExps    []*pingExp
	beginTxExps []*beginTxExp
	accountExps []*accountExp
	profileExps []*profileExp
}
type pingExp struct {
	err error
}

func (this *DB) ExpectPing(ctx context.Context, err error) {
	this.pingExps = append(this.pingExps, &pingExp{
		err: err,
	})
}

func (this *DB) Ping(context.Context) error {
	if len(this.pingExps) == 0 {
		return nil
	}

	exp := this.pingExps[0]

	this.pingExps = this.pingExps[1:]

	return exp.err
}

func (this *DB) ExpectBeginTx(ctx context.Context, err error) {
	this.beginTxExps = append(this.beginTxExps, &beginTxExp{
		err: err,
	})
}

func (this *DB) BeginTx(ctx context.Context) (db.ITx, error) {
	if len(this.beginTxExps) == 0 {
		return nil, nil
	}

	exp := this.beginTxExps[0]

	this.beginTxExps = this.beginTxExps[1:]

	return nil, exp.err
}

type accountExp struct {
	resp db.IAccountRepository
}

func (this *DB) ExpectAccount(ctx context.Context, resp db.IAccountRepository) {
	this.accountExps = append(this.accountExps, &accountExp{
		resp: resp,
	})
}

func (this *DB) Account(ctx context.Context) db.IAccountRepository {
	if len(this.accountExps) == 0 {
		return nil
	}

	exp := this.accountExps[0]

	this.accountExps = this.accountExps[1:]

	return exp.resp
}

type profileExp struct {
	resp db.IProfileRepository
}

func (this *DB) ExpectProfile(ctx context.Context, resp db.IProfileRepository) {
	this.profileExps = append(this.profileExps, &profileExp{
		resp: resp,
	})
}

func (this *DB) Profile(ctx context.Context) db.IProfileRepository {
	if len(this.profileExps) == 0 {
		return nil
	}

	exp := this.profileExps[0]

	this.profileExps = this.profileExps[1:]

	return exp.resp
}
