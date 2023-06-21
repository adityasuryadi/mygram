package usecase

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	domains "mygram/domains"
	entities "mygram/domains/entity"
	userEntities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/mail"
	"mygram/infrastructures/security"
	"mygram/infrastructures/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var mtx sync.Mutex

func NewUserUseCase(repository domains.UserRepository, userTokenRepo domains.UserTokenRepository, validate validation.Validation, database *gorm.DB, mailService mail.Mail) domains.UserUsecase {
	return &UserUseCaseImpl{
		db:            database,
		repository:    repository,
		userTokenRepo: userTokenRepo,
		Validate:      validate,
		MailService:   mailService,
	}
}

type UserUseCaseImpl struct {
	db            *gorm.DB
	repository    domains.UserRepository
	userTokenRepo domains.UserTokenRepository
	Validate      validation.Validation
	MailService   mail.Mail
}

// RegisterUser implements domains.UserUsecase
func (usecase *UserUseCaseImpl) RegisterUser(request model.RegisterUserRequest) (string, interface{}) {
	user := &userEntities.User{
		UserName: request.Username,
		Email:    request.Email,
		Password: security.GetHash([]byte(request.Password)),
		Age:      request.Age,
	}

	responseCode := make(chan string, 1)

	validationErr := usecase.Validate.ValidateRequest(request)
	if validationErr != nil {
		responseCode <- "400"
		return <-responseCode, validationErr
	}

	err := usecase.repository.Insert(user)
	if err != nil {
		return "500", nil
	}

	go func() {
		// sent mail
		usecase.MailService.SendMail(request.Email, "Hello <h1>"+request.Username+"</h1> you success register!")
	}()

	return "200", nil
}

func (usecase *UserUseCaseImpl) FetchUserLogin(request model.LoginUserRequest) (string, string) {
	email := request.Email
	result, err := usecase.repository.GetUserByEmail(email)

	now := time.Now()
	dateDiff := now.Sub(result.UserToken.ExpiredAt)

	// check if has user_token and not expired
	if !reflect.ValueOf(result.UserToken).IsZero() && dateDiff < 0 {
		return result.UserToken.Token, "200"
	}

	errorCode := make(chan string, 1)
	var token string
	token = ""
	if reflect.ValueOf(result).IsZero() {
		fmt.Println(err)
		errorCode <- "404"
	}

	if !reflect.ValueOf(result).IsZero() {
		hashPassword := result.Password
		err = security.ComparePassword(hashPassword, request.Password)
		if err != nil {
			errorCode <- "400"
		} else {
			token, _ = security.ClaimToken(email)
			errorCode <- "200"
		}
	}

	// store token into user token

	tx := usecase.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = usecase.userTokenRepo.InsertTokenWithTx(tx, result, token)
	if err != nil {
		errorCode <- "404"
	}

	tx.Commit()

	return token, <-errorCode
}

func (usecase *UserUseCaseImpl) UpdateUserRole(ctx *fiber.Ctx) (string, interface{}, interface{}) {
	responseCode := make(chan string, 1)
	id := ctx.Params("id")
	var request model.UpdateUserRoleRequest
	ctx.BodyParser(&request)
	var roles []entities.Role
	for _, v := range request.Roles {
		roles = append(roles, entities.Role{
			Id: int(v),
		})
	}
	responseCode <- "200"
	usecase.repository.AssignRole(id, roles)
	return <-responseCode, nil, nil
}

func (usecase *UserUseCaseImpl) Logout(ctx *fiber.Ctx) error {
	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)

	user, err := usecase.repository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	usecase.userTokenRepo.RemoveToken(user.Id.String())

	return nil
}
