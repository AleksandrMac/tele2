package storage

import (
	"context"
	"fmt"

	service "github.com/AleksandrMac/tele2/pkg/user/service"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type user struct {
	id   string
	name string
}

type Storage struct {
	db DBTX
}

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	// Prepare(ctx context.Context, name, sql string) (pgconn.StatementDescription, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func New(db DBTX) *Storage {
	return &Storage{db: db}
}

type StorageTx struct {
	*Storage
	db *pgxpool.Pool
}

func NewTx(db *pgxpool.Pool) *StorageTx {
	return &StorageTx{
		Storage: New(db),
		db:      db,
	}
}

func (x *StorageTx) execTx(ctx context.Context, f func(*Storage) error) error {
	tx, err := x.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = f(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

func (x *StorageTx) CreateTx(ctx context.Context, u *service.User) (id int64, err error) {
	err = x.execTx(ctx, func(s *Storage) error {
		id, err = s.Create(ctx, u)
		if err != nil {
			return err
		}
		return nil
	})
	return
}
func (x *Storage) Create(ctx context.Context, u *service.User) (id int64, err error) {
	return
}

func (x *Storage) Read(ctx context.Context, id int64) (*service.User, error) {
	return nil, nil
}

func (x *Storage) Update(ctx context.Context, id int64, u *service.User, fields ...string) error {
	return nil
}

func (x *Storage) Delete(ctx context.Context, id int64) error {
	return nil
}

func (x *Storage) GetList(ctx context.Context, limit, offset int64, filter string) ([]service.User, error) {
	return nil, nil
}
