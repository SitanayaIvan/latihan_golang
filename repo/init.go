package repo

import (
	ru "latihan_golang/repo/users"

	"gorm.io/gorm"
)

type Repo struct {
	User ru.RepoUserItf
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		User: ru.NewRepoUser(db),
	}
}
