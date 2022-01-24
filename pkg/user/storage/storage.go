package storage

import (
	"context"

	"github.com/AleksandrMac/tele2/pkg/user/model"
)

type Storage struct {
	DB CRUDL
}

type CRUDL interface {
	Create(ctx context.Context, u *model.UserDB) (id int64, err error)
	Read(ctx context.Context, id int64) (*model.UserDB, error)
	// Update обновляет указанные в fields список полей, если fields = nil, происходит обновление всех полей.
	Update(ctx context.Context, id int64, u *model.UserDB, fields ...string) error
	Delete(ctx context.Context, id int64) error
	GetList(ctx context.Context, limit, offset int64, filter string) ([]model.UserDB, error)
}

// type DB interface {
// 	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
// 	// Prepare(ctx context.Context, name, sql string) (pgconn.StatementDescription, error)
// 	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
// 	QueryRow(context.Context, string, ...interface{}) pgx.Row
// }
