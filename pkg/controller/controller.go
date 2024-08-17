package controller

import "github.com/Amiyaer/controller-formula/pkg/config"

type Controller struct {
    Config *config.Config
}

func New(cfg *config.Config) *Controller {
	return &Controller{
		Config: cfg,
	}
}

func (c *Controller) Run() error {
	// TODO
	return nil
}