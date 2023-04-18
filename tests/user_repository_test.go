package tests

import (
	"fmt"
	"log"
	entities "mygram/domains/entity"
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


func TestCreateRole(t *testing.T) {
	db:=initDB()
	roleRepository:=repository.NewRoleRepository(db)
	request := entities.Role{
		Name: "user",
		Permissions: []entities.Permission{
			{Id: 1},
			{Id: 2},
			{Id: 3},
			{Id: 5},
		},
	}
	err := roleRepository.Insert(&request)
	if err != nil {
		log.Print(err)
	}
}


func TestGetRole(t *testing.T){
	db:=initDB()
	var role entities.Role
	db.Where("name = ? ","admin").Preload("Permissions").First(&role)
	found:=false
	// var tmp map[int]string
	// tmp := map[int]string{}
	// tmp := []string{}
	for _, v := range role.Permissions {
		// tmp = append(tmp,v.Name)
		// tmp[i] = v.Name
		if v.Name == "permission.create" {
			found = true
			break
		}
	}
	log.Print(found)
	
}

func TestAssignRole(t *testing.T){
	db:=initDB()
	userRepository := repository.NewUserRepositoryPostgres(db)
	roles := []int{
		1,2,
	}
	userRepository.AssignRole("33a61a3d-88e8-484d-8061-3db0bff92e3a",roles)
}

func TestGetUserRole(t *testing.T){
	db:=initDB()
	var user entities.User
	// var roles []string
	permissions := []string{}
	db.Preload("Roles.Permissions").Where("id = ?","33a61a3d-88e8-484d-8061-3db0bff92e3a").First(&user)
	// log.Print(user.Roles)
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			permissions = append(permissions, permission.Name)
		}
	}
	log.Println(permissions)
}

// func findInSlice(value interface{}){
	
// }