package tests

import (
	"fmt"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	"testing"

	"gorm.io/gorm"
)

func initDB() *gorm.DB{
	db:=dbConfig.NewTestPostgresDB()
	return db	
}

func TestGetUserByEmail(t *testing.T) {
	db:= initDB()
	email := "aditsss@mail.com"
	userRepository := repository.NewUserRepositoryPostgres(db)
	user,err := userRepository.GetUserByEmail(email)
	if err != nil {
		fmt.Println(user)
		return
	}
	fmt.Println(user)
}
