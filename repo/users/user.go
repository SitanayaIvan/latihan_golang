package repo_users

import (
	dm "github.com/SitanayaIvan/latihan_golang/domains"
	"gorm.io/gorm"
)

type RepoUserItf interface {
	CreateUser(userProfile dm.UserProfile) error
	GetUsers() ([]dm.UserProfile, error)
	GetUserById(id int) (dm.UserProfile, error)
	UpdateUser(userProfile dm.UserProfile, id int) error
	DeleteUser(id int) error
}

type RepoUser struct {
	Db *gorm.DB
}

func NewRepoUser(db *gorm.DB) *RepoUser {
	return &RepoUser{db}
}

func (ru RepoUser) CreateUser(userProfile dm.UserProfile) error {
	err := ru.Db.Create(&userProfile).Error
	return err
}

func (ru RepoUser) GetUsers() ([]dm.UserProfile, error) {
	var userProfiles []dm.UserProfile

	err := ru.Db.Find(&userProfiles).Error
	return userProfiles, err
}

func (ru RepoUser) GetUserById(id int) (dm.UserProfile, error) {
	var userProfile dm.UserProfile

	err := ru.Db.First(&userProfile, id).Error
	return userProfile, err
}

func (ru RepoUser) UpdateUser(userProfile dm.UserProfile, id int) error {
	// check user by id
	err := ru.Db.First(&dm.UserProfile{}, id).Error
	if err != nil {
		return err
	}

	// update user profile
	err = ru.Db.Model(&dm.UserProfile{}).Where("user_profiles.id=?", id).
		Updates(dm.UserProfile{
			FirstName: userProfile.FirstName,
			LastName:  userProfile.LastName,
			Age:       userProfile.Age,
			Email:     userProfile.Email,
			Password:  userProfile.Password,
		}).Error
	return err
}

func (ru RepoUser) DeleteUser(id int) error {
	// check user by id
	err := ru.Db.First(&dm.UserProfile{}, id).Error
	if err != nil {
		return err
	}

	// delete user profile
	err = ru.Db.Delete(&dm.UserProfile{}, id).Error
	return err
}
