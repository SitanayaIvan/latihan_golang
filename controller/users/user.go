package controller_users

import (
	co "latihan_golang/constants"
	dm "latihan_golang/domains"
	ru "latihan_golang/repo/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllerUser struct {
	User ru.RepoUserItf
}

func NewControllerUser(user ru.RepoUserItf) *ControllerUser {
	return &ControllerUser{
		User: user,
	}
}

var httpStatus int

func (cu *ControllerUser) CreateUser(c *gin.Context) {
	var user dm.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		httpStatus = 400
		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	err = cu.User.CreateUser(user)
	if err != nil {
		httpStatus = 500
		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	httpStatus = 200
	c.JSON(httpStatus, dm.Success{
		Status: co.StatusSuccess,
	})
}

func (cu *ControllerUser) GetUsers(c *gin.Context) {
	searching := c.Query("searching")
	sorting := c.Query("sorting")
	filterRole := c.Query("filter-role")

	users, countRow, err := cu.User.GetUsers(searching, sorting, filterRole)
	if err != nil {
		httpStatus = 400
		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	httpStatus = 200
	c.JSON(httpStatus, dm.Success{
		Status: co.StatusSuccess,
		Data:   users,
		Count:  countRow,
	})
}

func (cu *ControllerUser) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := cu.User.GetUserById(uint(id))
	if err != nil {
		if err.Error() == co.IdNotFound {
			httpStatus = 400
		} else {
			httpStatus = 500
		}

		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	httpStatus = 200
	c.JSON(httpStatus, dm.Success{
		Status: co.StatusSuccess,
		Data:   user,
	})
}

func (cu *ControllerUser) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user dm.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		httpStatus = 400
		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	err = cu.User.UpdateUser(uint(id), user)
	if err != nil {
		if err.Error() == co.IdNotFound {
			httpStatus = 400
		} else {
			httpStatus = 500
		}

		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	httpStatus = 200
	c.JSON(httpStatus, dm.Success{
		Status: co.StatusSuccess,
	})
}

func (cu *ControllerUser) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := cu.User.DeleteUser(uint(id))
	if err != nil {
		if err.Error() == co.IdNotFound {
			httpStatus = 400
		} else {
			httpStatus = 500
		}

		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	httpStatus = 200
	c.JSON(httpStatus, dm.Success{
		Status: co.StatusSuccess,
	})
}
