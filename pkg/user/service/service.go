package service

import (
	"context"

	"github.com/AleksandrMac/tele2/pkg/user/model"
	"github.com/AleksandrMac/tele2/pkg/user/storage"
)

type Services interface {
	Create(ctx context.Context, name string) (id string, err error)
	Get(ctx context.Context, id uint64) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, userID uint64) error
	DeleteByEmail(ctx context.Context, email string) error
	FindUserByName(ctx context.Context, name string, limit, offset int64) ([]model.User, error)
}

type Service struct {
	storage storage.CRUDL
}

func New(storage storage.CRUDL) *Service {
	return &Service{
		storage: storage,
	}
}

func (x *Service) Create(ctx context.Context, name string) (id string, err error) {
	return "", nil
}

func (x *Service) Get(ctx context.Context, id uint64) (*model.User, error) {
	return nil, nil
}

func (x *Service) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

func (x *Service) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (x *Service) Delete(ctx context.Context, userID uint64) error {
	return nil
}

func (x *Service) DeleteByEmail(ctx context.Context, email string) error {
	return nil
}

func (x *Service) FindUserByName(ctx context.Context, name string, limit, offset int64) ([]model.User, error) {
	return nil, nil
}
