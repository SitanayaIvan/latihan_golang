package repo_users

import (
	"errors"
	co "latihan_golang/constants"
	dm "latihan_golang/domains"
	"strings"

	"gorm.io/gorm"
)

type RepoUserItf interface {
	CreateUser(user dm.User) error
	GetUsers(searching string, sorting string, filterRole string) ([]dm.User, int64, error)
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

func (ru *RepoUser) CreateUser(user dm.User) error {

	// validate email
	isEmailFound := ru.Db.Where("lower(users.email) = ?", strings.ToLower(user.Email)).
		First(&dm.User{}).RowsAffected
	if isEmailFound > 0 {
		return errors.New(co.EmailExists)
	}

	err := ru.Db.Create(&user).Error
	return err
}

func (ru *RepoUser) GetUsers(searching string, sorting string, filterRole string) ([]dm.User, int64, error) {
	var users []dm.User
	var countRow int64
	var err error

	searching = "%" + strings.ToLower(searching) + "%"
	filterRole = strings.ToLower(filterRole)
	sorting = strings.ToLower(sorting)

	// filter role
	if len(strings.TrimSpace(filterRole)) > 0 {
		err = ru.Db.Model(&dm.User{}).
			Where("(lower(users.first_name) like ? or lower(users.last_name) like ? or lower(users.email) like ?) and lower(users.role)=?", searching, searching, searching, filterRole).
			Count(&countRow).Order(sorting).Find(&users).Error
		return users, countRow, err
	}

	err = ru.Db.Model(&dm.User{}).
		Where("lower(users.first_name) like ? or lower(users.last_name) like ? or lower(users.email) like ?", searching, searching, searching).
		Count(&countRow).Order(sorting).Find(&users).Error
	return users, countRow, err
}

func (ru *RepoUser) GetUserById(id uint) (dm.User, error) {
	var user dm.User

	err := ru.Db.First(&user, id).Error
	if err != nil {
		return user, errors.New(co.IdNotFound)
	}
	return user, err
}

func (ru *RepoUser) UpdateUser(id uint, user dm.User) error {
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

func (ru *RepoUser) DeleteUser(id uint) error {
	// search for id
	err := ru.Db.First(&dm.User{}, id).Error
	if err != nil {
		return errors.New(co.IdNotFound)
	}

	// delete user
	err = ru.Db.Delete(&dm.User{}, id).Error
	return err
}
