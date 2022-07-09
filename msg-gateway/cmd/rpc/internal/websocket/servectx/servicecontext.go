package servectx

import (
	"work-test/msg-gateway/cmd/rpc/internal/websocket/config"
)

type ServiceContext struct {
	Config      config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
