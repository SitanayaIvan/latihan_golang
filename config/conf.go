package config

import (
	"log"
	"os"

	dm "github.com/SitanayaIvan/latihan_golang/domains"
	"github.com/alexsasharegan/dotenv"
)

func LoadEnv() dm.ConfigBody {
	var conf dm.ConfigBody

	err := dotenv.Load(".env")
	if err != nil {
		log.Println("[ERROR] load .env failed")
	}

	// App
	conf.App.Port = os.Getenv("APP_PORT")
	conf.App.Env = os.Getenv("APP_ENV")

	// DB
	conf.Db.Host = os.Getenv("DB_HOST")
	conf.Db.User = os.Getenv("DB_USER")
	conf.Db.Password = os.Getenv("DB_PASSWORD")
	conf.Db.Name = os.Getenv("DB_NAME")
	conf.Db.Port = os.Getenv("DB_PORT")

	return conf
}
