package tests

import (
	"encoding/json"
	"io/ioutil"
	"mygram/applications/usecase"
	"mygram/commons/exceptions"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	"mygram/infrastructures/validation"
	handler "mygram/interfaces/http/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var jwtToken string

// var userId string

func TestCreateUser(t *testing.T) {
	app:=fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"username": "adit",
		"email": "adit@mail.com",
		"password": "admin",
		"age": 25
	  }`)

	request := httptest.NewRequest(http.MethodPost, "user", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].(map[string]interface{})
	parse := response
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "OK", parse["status"])
	assert.Equal(t, float64(200), parse["code"])
	assert.Equal(t, data["name"], "Aditya")
	assert.Equal(t, data["age"], "25")
	assert.Equal(t, data["email"], "adit@mail.com")
}

/**
* Test Create User
 */

func TestCreateEmptyEmail(t *testing.T) {
	app:=fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"username": "adit",
		"email": "",
		"password": "admin",
		"age": 25
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/user", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].([]interface{})
	parse := response
	assert.Equal(t, 400, res.StatusCode)
	assert.Equal(t, "BAD_REQUEST", parse["status"])
	assert.Equal(t, float64(400), parse["code"])

	for _, val := range data {
		value := val.(map[string]interface{})
		assert.Equal(t, "email", value["field"])
		assert.Equal(t, "field tidak boleh kosong", value["message"])
	}
}

func TestCreateEmptyPassword(t *testing.T) {
	app:=fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"username": "adit",
		"email": "adit@mail.com",
		"password": "",
		"age": 25
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/user", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].([]interface{})
	parse := response
	assert.Equal(t, 400, res.StatusCode)
	assert.Equal(t, "BAD_REQUEST", parse["status"])
	assert.Equal(t, float64(400), parse["code"])

	for _, val := range data {
		value := val.(map[string]interface{})
		assert.Equal(t, "password", value["field"])
		assert.Equal(t, "field tidak boleh kosong", value["message"])
	}
}

func TestCreateEmptyAge(t *testing.T) {
	app:=fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"username": "adit",
		"email": "adit@mail.com",
		"password": "1234567",
		"age": 
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/user", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].([]interface{})
	parse := response
	assert.Equal(t, 400, res.StatusCode)
	assert.Equal(t, "BAD_REQUEST", parse["status"])
	assert.Equal(t, float64(400), parse["code"])

	for _, val := range data {
		value := val.(map[string]interface{})
		assert.Equal(t, "password", value["field"])
		assert.Equal(t, "field tidak boleh kosong", value["message"])
	}
}

func TestCreateWrongEmail(t *testing.T) {
	app:=fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"username": "adit",
		"email": "adit@mail",
		"password": "admin",
		"age": 25
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/user", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].([]interface{})
	parse := response
	assert.Equal(t, 400, res.StatusCode)
	assert.Equal(t, "BAD_REQUEST", parse["status"])
	assert.Equal(t, float64(400), parse["code"])

	for _, val := range data {
		value := val.(map[string]interface{})
		assert.Equal(t, "email", value["field"])
		assert.Equal(t, "format email salah", value["message"])
	}
}

func Setup() handler.UserHandler {
	configApp := config.New()
	db:=dbConfig.NewPostgresDB(configApp)
	validate:=validation.NewValidation(db)

	// user
	userRepository:=repository.NewUserRepositoryPostgres(db)
	userUsecase:=usecase.NewUserUseCase(userRepository,validate)
	userHandler:=handler.NewUserHandler(userUsecase)
	return userHandler
}