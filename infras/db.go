package infras

import (
	dm "latihan_golang/domains"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(conf dm.DbAccount) (*gorm.DB, error) {
	dsn := "host=" + conf.Host +
		" user=" + conf.User +
		" password=" + conf.Password +
		" dbname=" + conf.Name +
		" port=" + conf.Port +
		" sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func MigrateTable(db *gorm.DB) {
	db.AutoMigrate(&dm.User{})
}
