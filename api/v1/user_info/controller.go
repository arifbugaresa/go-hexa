package user_info

import "github.com/arifbugaresa/go-hexa/business/user_info"

type Controller struct {
	service user_info.Service
}

func NewController(service user_info.Service) *Controller {
	return &Controller{
		service: service,
	}
}

