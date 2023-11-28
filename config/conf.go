package config

import (
	dm "latihan_golang/domains"
	"os"

	"github.com/alexsasharegan/dotenv"
)

func LoadEnv() (dm.ConfigBody, error) {
	var confBody dm.ConfigBody

	err := dotenv.Load(".env")
	if err != nil {
		err = nil
	}

	// app
	confBody.App.Port = os.Getenv("APP_PORT")
	confBody.App.Env = os.Getenv("APP_ENV")

	// Db
	confBody.Db.Host = os.Getenv("DB_HOST")
	confBody.Db.User = os.Getenv("DB_USER")
	confBody.Db.Password = os.Getenv("DB_PASSWORD")
	confBody.Db.Name = os.Getenv("DB_NAME")
	confBody.Db.Port = os.Getenv("DB_PORT")

	return confBody, err
}
