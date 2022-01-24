package http

import (
	"github.com/AleksandrMac/tele2/pkg/user/service"
	"github.com/rs/zerolog"
)

type HTTP struct {
	Log *zerolog.Logger
	service.Services
}

func New(log *zerolog.Logger, service service.Services) *HTTP {
	return &HTTP{
		Log:      log,
		Services: service,
	}
}
