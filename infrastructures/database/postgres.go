package database

import (
	"mygram/commons/exceptions"
	config "mygram/infrastructures"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(configuration config.Config) *gorm.DB {
	host := configuration.Get("POSTGRE_HOST")
	user := configuration.Get("POSTGRE_USER")
	password := configuration.Get("POSTGRE_PASSWORD")
	port := configuration.Get("POSTGRE_PORT")
	db_name := configuration.Get("POSTGRE_DB_NAME")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + db_name + " port=" + port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exceptions.PanicIfNeeded(err)

	return db
}
