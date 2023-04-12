package middleware

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"

	// repository "mygram/infrastructures/repository/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Can(ctx *fiber.Ctx,permission string) error {
	configApp := config.New()
	db:=dbConfig.NewPostgresDB(configApp)
	
	var user entities.User

	usertoken := ctx.Locals("user").(*jwt.Token)
	claims := usertoken.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	err := db.Where("email",email).First(&user).Error

	if err != nil {
		panic(err)
	}

	// var roles []string
	// permissions := []string{}
	found:=false
	db.Preload("Roles.Permissions").Where("id = ?","33a61a3d-88e8-484d-8061-3db0bff92e3a").First(&user)
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			// permissions = append(permissions, permission.Name)
			if permission.Name == "permission.create" {
				found = true
				break
			}
		}
	}
	
	if found {
		return ctx.Next()
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
		Code:   fiber.StatusUnauthorized,
		Status: "UNAUTHORIZE",
		Message: "UNAUTHORIZE",
		Data:   nil,
	})
}