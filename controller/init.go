package controller

import (
	cu "github.com/SitanayaIvan/latihan_golang/controller/users"
	r "github.com/SitanayaIvan/latihan_golang/repo"
)

type Controller struct {
	User cu.ControllerUser
}

func NewController(repo r.Repo) *Controller {
	return &Controller{
		User: *cu.NewControllerUser(repo.User),
	}
}
