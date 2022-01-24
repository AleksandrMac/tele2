package grpc

import (
	"github.com/AleksandrMac/tele2/pkg/user/service"
	"github.com/rs/zerolog"
)

type GRPC struct {
	Log *zerolog.Logger
	service.Services
}

func New(log *zerolog.Logger, service service.Services) *GRPC {
	return &GRPC{
		Log:      log,
		Services: service,
	}
}
