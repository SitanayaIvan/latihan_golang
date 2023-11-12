package main

import (
	"latihan_golang/config"
	"latihan_golang/controller"
	"latihan_golang/infras"
	repo "latihan_golang/repo"
	"latihan_golang/routes"
	"log"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {

	// Load env
	conf, err := config.LoadEnv()
	if err != nil {
		log.Println("[ERROR] load .env failed!")
	}

	log.Println("[GIN ENGINE] Starting...")
	switch conf.App.Env {
	case "production", "staging":
		gin.SetMode("release")
	default:
		gin.SetMode("debug")
	}

	// Db
	db, err := infras.ConnectDb(conf.Db)
	if err != nil {
		log.Fatal("[ERROR] connect to db failed!")
	}
	log.Println("[Database] Connected")

	// migrate table
	// infras.MigrateTable(db)

	// routes
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(cors.AllowAll())

	repo := repo.NewRepo(db)
	controller := controller.NewController(*repo)

	// endpoints
	routes.GetEndpoints(r, controller)

	log.Println("[SERVER] Started")
	r.Run(":" + conf.App.Port)
}
