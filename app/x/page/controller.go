package page

import "api/common"

type InjectController struct {
	*common.App
	Service *Service
}

type Controller struct {
	*InjectController
}

func NewController(i *InjectController) *Controller {
	return &Controller{
		InjectController: i,
	}
}