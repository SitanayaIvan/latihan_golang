package main

import (
	"log"

	"github.com/SitanayaIvan/latihan_golang/config"
	ctl "github.com/SitanayaIvan/latihan_golang/controller"
	"github.com/SitanayaIvan/latihan_golang/infras"
	re "github.com/SitanayaIvan/latihan_golang/repo"
	"github.com/SitanayaIvan/latihan_golang/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	conf := config.LoadEnv()

	// Switch gin mode
	log.Println("[GIN ENGINE] STARTING...")
	switch conf.App.Env {
	case "production", "staging":
		gin.SetMode("release")
	default:
		gin.SetMode("debug")
	}

	// Connect to DB
	db, err := infras.ConnectDb(conf.Db)
	if err != nil {
		log.Fatal("[ERROR] connect to db failed!")
	}
	log.Println("[DATABASE] CONNECTED!")

	// Migrate table
	// infras.MigrateTable(db)

	// gin engine
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(cors.AllowAll())

	repo := re.NewRepo(db)
	controller := ctl.NewController(*repo)

	// route
	routes.GetEndpoint(r, *controller)

	log.Println("[SERVER] STARTED!")
	r.Run(":" + conf.App.Port)
}
