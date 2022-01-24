package pgxdb

import (
	"context"

	"github.com/AleksandrMac/tele2/pkg/user/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PGXDB struct {
	db *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *PGXDB {
	return &PGXDB{
		db: pool,
	}
}

func (x *PGXDB) Create(ctx context.Context, u *model.UserDB) (id int64, err error) {
	return
}

func (x *PGXDB) Read(ctx context.Context, id int64) (*model.UserDB, error) {
	return nil, nil
}

func (x *PGXDB) Update(ctx context.Context, id int64, u *model.UserDB, fields ...string) error {
	return nil
}

func (x *PGXDB) Delete(ctx context.Context, id int64) error {
	return nil
}

func (x *PGXDB) GetList(ctx context.Context, limit, offset int64, filter string) ([]model.UserDB, error) {
	return nil, nil
}
