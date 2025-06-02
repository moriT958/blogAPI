package controllers

import "github.com/moriT958/go-api/services"

type Controller struct {
	Service services.IBlogService
}

func NewController(s services.IBlogService) *Controller {
	return &Controller{
		Service: s,
	}
}
