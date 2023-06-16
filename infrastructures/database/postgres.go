package database

import (
	"time"

	"mygram/commons/exceptions"
	config "mygram/infrastructures"

	customlog "mygram/infrastructures/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(configuration config.Config) *gorm.DB {
	host := configuration.Get("POSTGRE_HOST")
	user := configuration.Get("POSTGRE_USER")
	password := configuration.Get("POSTGRE_PASSWORD")
	port := configuration.Get("POSTGRE_PORT")
	db_name := configuration.Get("POSTGRE_DB_NAME")

	newLogger := logger.New(
		customlog.NewLog(),
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + db_name + " port=" + port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	exceptions.PanicIfNeeded(err)

	return db
}

func NewTestPostgresDB() *gorm.DB {
	host := "localhost"
	user := "postgres"
	password := "postgres"
	port := "5433"
	db_name := "mygram"

	newLogger := logger.New(
		customlog.NewLog(),
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + db_name + " port=" + port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	exceptions.PanicIfNeeded(err)

	return db
}
