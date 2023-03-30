package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/infrastructures/validation"
	"mygram/interfaces/http/api/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewPhotoHandler(usecase domains.PhotoUsecase,validate validator.Validate) PhotoHandler {
	return PhotoHandler{
		usecase: usecase,
		validate: validate,
	}
}

type PhotoHandler struct {
	usecase domains.PhotoUsecase
	validate validator.Validate
}

func (handler PhotoHandler) Route(app *fiber.App){
	app.Post("photo",middleware.Verify(),handler.PostPhoto)	
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func (handler PhotoHandler) PostPhoto(ctx *fiber.Ctx) error {
	var request model.CreatePhotoRequest
	ctx.BodyParser(&request)
	err := handler.validate.Struct(&request)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	
	request.Email = email
	// headers := ctx.GetReqHeaders()
	// for i, _ := range headers {
	// 	fmt.Println(i)
	// }
	// fmt.Println(ctx.GetReqHeaders())
	
	// token, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("AllYourBase"), nil
	// })
	
	// if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
	// 	fmt.Printf("%v %v", claims.Foo, claims.RegisteredClaims.Issuer)
	// } else {
	// 	fmt.Println(err)
	// }

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		out := make([]validation.ErrorMessage,len(validationErrors))
		for i, fieldError := range validationErrors {
			out[i] = validation.ErrorMessage{
				fieldError.Field(),
				validation.GetErrorMsg(fieldError),
			}
		}
		model.BadRequestResponse(ctx,"CLIENT SERVER ERROR",out)
		return nil
	}

	_,errCode := handler.usecase.CreatePhoto(request)
	if errCode == "200"{
		model.SuccessResponse(ctx,"SUCCESS CREATE PHOTO",nil)
	}

	if errCode == "500" {
		model.InternalServerErrorResponse(ctx,"INTERNAL SERVER ERROR",nil)
	}
	return nil
}