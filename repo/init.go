package repo

import (
	ru "github.com/SitanayaIvan/latihan_golang/repo/users"
	"gorm.io/gorm"
)

type Repo struct {
	User ru.RepoUser
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		User: *ru.NewRepoUser(db),
	}
}
