package helper

import (
	entities "mygram/domains/entity"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"
	// repository "mygram/infrastructures/repository/postgres"
)

func Can(email string, permissionName string) bool {
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
			// permissions = append(permissions, permission.Name)
			if permission.Name == permissionName {
				found = true
				break
			}
		}
	}

	return found
}