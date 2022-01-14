package service

import (
	"context"
	"fmt"
	"strconv"

	http "github.com/AleksandrMac/tele2/pkg/user/http"
)

type User struct {
	ID   int64
	Name string
}

type CRUDL interface {
	Create(ctx context.Context, u *User) (id int64, err error)
	Read(ctx context.Context, id int64) (*User, error)
	// Update обновляет указанные в fields список полей, если fields = nil, происходит обновление всех полей.
	Update(ctx context.Context, id int64, u *User, fields ...string) error
	Delete(ctx context.Context, id int64) error
	GetList(ctx context.Context, limit, offset int64, filter string) ([]User, error)
}

type Service struct {
	storage CRUDL
}

func New(storage CRUDL) *Service {
	return &Service{
		storage: storage,
	}
}

func (x *Service) CreateUser(ctx context.Context, name string) (id string, err error) {
	uid, err := x.storage.Create(ctx, &User{
		Name: name,
	})
	return strconv.Itoa(int(uid)), err
}

func (x *Service) FindUserByName(ctx context.Context, name string, limit, offset int64) ([]http.User, error) {
	u, err := x.storage.GetList(ctx, limit, offset, fmt.Sprintf(`name LIKE '%s%%'`, name))
	if err != nil {
		return nil, fmt.Errorf(`service-find-user-by-name: %w`, err)
	}
	outUser := make([]http.User, len(u))
	for i := range u {
		outUser[i] = http.User{
			ID:   strconv.Itoa(int(u[i].ID)),
			Name: u[i].Name,
		}
	}
	return outUser, nil
}
