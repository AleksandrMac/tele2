package graphql

import (
	"github.com/AleksandrMac/tele2/pkg/user/service"
	"github.com/rs/zerolog"
)

type GRAPHQL struct {
	Log *zerolog.Logger
	service.Services
}

func New(log *zerolog.Logger, service service.Services) *GRAPHQL {
	return &GRAPHQL{
		Log:      log,
		Services: service,
	}
}
