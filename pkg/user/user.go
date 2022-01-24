package user

import (
	"github.com/AleksandrMac/tele2/pkg/user/http"
	"github.com/AleksandrMac/tele2/pkg/user/service"
	"github.com/AleksandrMac/tele2/pkg/user/storage"
	"github.com/AleksandrMac/tele2/pkg/user/storage/pgxdb"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

// аналогично с NewHTTP создаются хэндлеры для graphql и grpc
func NewHTTP(log *zerolog.Logger, service service.Services) *http.HTTP {
	return http.New(log, service)
}

func NewService(storage storage.CRUDL) service.Services {
	return service.New(storage)
}

func NewStoragePGX(dbpool *pgxpool.Pool) storage.CRUDL {
	return pgxdb.New(dbpool)
}
