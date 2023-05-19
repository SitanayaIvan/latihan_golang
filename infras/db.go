package infras

import (
	dm "github.com/SitanayaIvan/latihan_golang/domains"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb(conf dm.DBAccount) (*gorm.DB, error) {

	dsn := "host=" + conf.Host +
		" user=" + conf.User +
		" password=" + conf.Password +
		" dbname=" + conf.Name +
		" port=" + conf.Port +
		" sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Error)})

	return db, err
}

func MigrateTable(db *gorm.DB) {
	db.AutoMigrate(&dm.UserProfile{})
}
