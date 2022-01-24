package user

import (
	"github.com/AleksandrMac/tele2/pkg/user/http"
	"github.com/AleksandrMac/tele2/pkg/user/service"
	"github.com/AleksandrMac/tele2/pkg/user/storage"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

// аналогично с NewHTTP создаются хэндлеры для graphql и grpc
func NewHTTP(log *zerolog.Logger, service http.Services) *http.HTTP {
	return http.New(log, service)
}

func NewService(storage service.CRUDL) *service.Service {
	return service.New(storage)
}

func NewStorage(db *pgxpool.Pool) *storage.StorageTx {
	return storage.NewTx(db)
}
