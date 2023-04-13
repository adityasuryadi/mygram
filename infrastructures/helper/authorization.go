package helper

import (
	entities "mygram/domains/entity"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"

	"github.com/google/uuid"
)

func Can(email string, permissionName string) (bool,uuid.UUID) {
	configApp := config.New()
	db := dbConfig.NewPostgresDB(configApp)

	var user entities.User
	err := db.Where("email", email).First(&user).Error

	if err != nil {
		panic(err)
	}

	found := false
	db.Preload("Roles.Permissions").First(&user)
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			if permission.Name == permissionName {
				found = true
				break
			}
		}
	}

	return found,user.Id
}