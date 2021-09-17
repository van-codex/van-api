package index

import (
	"github.com/kainonly/go-bit/authx"
	"github.com/kainonly/go-bit/crud"
	"go.uber.org/fx"
	"lab-api/common"
)

var Provides = fx.Provide(
	NewController,
	NewService,
)

type Controller struct {
	*ControllerInject
	*crud.API
	Auth *authx.Auth
}

type ControllerInject struct {
	common.App

	Service *Service
}

func NewController(i ControllerInject) *Controller {
	return &Controller{
		ControllerInject: &i,
		Auth:             i.Authx.Make("system"),
	}
}

type Service struct {
	*ServiceInject
	Key string
}

type ServiceInject struct {
	common.App
}

func NewService(i ServiceInject) *Service {
	return &Service{
		ServiceInject: &i,
		Key:           i.Set.RedisKey("code:"),
	}
}
