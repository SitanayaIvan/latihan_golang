package controller_users

import (
	"strconv"

	co "github.com/SitanayaIvan/latihan_golang/constants"
	dm "github.com/SitanayaIvan/latihan_golang/domains"
	ru "github.com/SitanayaIvan/latihan_golang/repo/users"
	"github.com/gin-gonic/gin"
)

type ControllerUser struct {
	User ru.RepoUserItf
}

func NewControllerUser(user ru.RepoUserItf) *ControllerUser {
	return &ControllerUser{user}
}

var httpStatus int

func (cu ControllerUser) CreateUser(c *gin.Context) {
	var userProfile dm.UserProfile

	err := c.ShouldBindJSON(&userProfile)
	if err != nil {
		httpStatus = 400
		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	err = cu.User.CreateUser(userProfile)
	if err != nil {
		if err.Error() == co.Unauthorized {
			httpStatus = 401
		} else if err.Error() == co.IdNotFound {
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
	c.JSON(httpStatus, gin.H{
		"status": co.StatusSuccess,
	})
}

func (cu ControllerUser) GetUsers(c *gin.Context) {
	userProfiles, err := cu.User.GetUsers()
	if err != nil {
		if err.Error() == co.Unauthorized {
			httpStatus = 401
		} else if err.Error() == co.IdNotFound {
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
	c.JSON(httpStatus, dm.GetApi{
		Status: co.StatusSuccess,
		Data:   userProfiles,
	})
}

func (cu ControllerUser) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	userprofile, err := cu.User.GetUserById(id)
	if err != nil {
		if err.Error() == co.Unauthorized {
			httpStatus = 401
		} else if err.Error() == co.IdNotFound {
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
	c.JSON(httpStatus, dm.GetApi{
		Status: co.StatusSuccess,
		Data:   userprofile,
	})

}

func (cu ControllerUser) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var userProfile dm.UserProfile
	err := c.ShouldBindJSON(&userProfile)
	if err != nil {
		httpStatus = 400
		c.JSON(httpStatus, dm.Error{
			Status:  co.StatusError,
			Message: err.Error(),
		})
		return
	}

	err = cu.User.UpdateUser(userProfile, id)
	if err != nil {
		if err.Error() == co.Unauthorized {
			httpStatus = 401
		} else if err.Error() == co.IdNotFound {
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
	c.JSON(httpStatus, gin.H{
		"status": co.StatusSuccess,
	})
}

func (cu ControllerUser) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := cu.User.DeleteUser(id)
	if err != nil {
		if err.Error() == co.Unauthorized {
			httpStatus = 401
		} else if err.Error() == co.IdNotFound {
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
	c.JSON(httpStatus, gin.H{
		"status": co.StatusSuccess,
	})
}
