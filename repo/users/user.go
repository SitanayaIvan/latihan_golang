package repo_users

import (
	"errors"
	co "latihan_golang/constants"
	dm "latihan_golang/domains"

	"gorm.io/gorm"
)

type RepoUserItf interface {
	CreateUser(user dm.User) error
	GetUsers() ([]dm.User, int64, error)
	GetUserById(id uint) (dm.User, error)
	UpdateUser(id uint, user dm.User) error
	DeleteUser(id uint) error
}

type RepoUser struct {
	Db *gorm.DB
}

func NewRepoUser(db *gorm.DB) *RepoUser {
	return &RepoUser{Db: db}
}

func (ru RepoUser) CreateUser(user dm.User) error {
	err := ru.Db.Create(&user).Error
	return err
}

func (ru RepoUser) GetUsers() ([]dm.User, int64, error) {
	var users []dm.User
	var countRow int64

	err := ru.Db.Model(&dm.User{}).Count(&countRow).Find(&users).Error
	return users, countRow, err
}

func (ru RepoUser) GetUserById(id uint) (dm.User, error) {
	var user dm.User

	err := ru.Db.First(&user, id).Error
	if err != nil {
		return user, errors.New(co.IdNotFound)
	}
	return user, err
}

func (ru RepoUser) UpdateUser(id uint, user dm.User) error {
	// search for id
	err := ru.Db.First(&dm.User{}, id).Error
	if err != nil {
		return errors.New(co.IdNotFound)
	}

	// update user
	err = ru.Db.Where("users.id=?", id).Updates(&dm.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Email:     user.Email,
		Password:  user.Password,
	}).Error
	return err
}

func (ru RepoUser) DeleteUser(id uint) error {
	// search for id
	err := ru.Db.First(&dm.User{}, id).Error
	if err != nil {
		return errors.New(co.IdNotFound)
	}

	// delete user
	err = ru.Db.Delete(&dm.User{}, id).Error
	return err
}
