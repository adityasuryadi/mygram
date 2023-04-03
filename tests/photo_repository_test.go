package tests

import (
	"log"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	"testing"
)

func TestFindPhotoById(t *testing.T) {
	db:=dbConfig.NewTestPostgresDB()
	id := "aditsss@mail.com"
	photoRepository := repository.NewPhotoRepository(db)
	result,err := photoRepository.FindById(id)
	if err != nil {
		log.Print(result)
	}
	log.Print(result)
}
