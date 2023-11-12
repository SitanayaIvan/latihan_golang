package controller

import (
	cu "latihan_golang/controller/users"
	"latihan_golang/repo"
)

type Controller struct {
	User cu.ControllerUser
}

func NewController(r repo.Repo) *Controller {
	return &Controller{
		User: *cu.NewControllerUser(r.User),
	}
}
