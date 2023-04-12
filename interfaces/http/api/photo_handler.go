package handler

import (
	"log"
	"mygram/domains"
	"mygram/domains/model"
	"mygram/infrastructures/helper"
	"mygram/infrastructures/security"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewPhotoHandler(usecase domains.PhotoUsecase) PhotoHandler {
	return PhotoHandler{
		usecase: usecase,
	}
}

type PhotoHandler struct {
	usecase domains.PhotoUsecase
}

func (handler PhotoHandler) Route(app *fiber.App){	
	// auth
	photo:=app.Group("photo",middleware.Verify())
	photo.Post("/",handler.PostPhoto)	
	photo.Get("/",handler.ListPhoto)
	photo.Get("/:id",handler.GetPhoto)
	photo.Put("/:id",handler.UpdatePhoto)
	photo.Delete("/:id",handler.DeletePhoto)
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func (handler PhotoHandler) PostPhoto(ctx *fiber.Ctx) error {
	var request model.CreatePhotoRequest
	ctx.BodyParser(&request)
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

	result,errCode := handler.usecase.CreatePhoto(request)
	model.GetResponse(ctx,errCode,"",result)
	return nil
}


// Registeruser List Photo 
// @Summary List Photo
// @Description List Photo
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.PhotoResponse}
// @Failure 400 {string} model.WebResponse{code=400}
// @Failure 500 {string} model.WebResponse{code=500}
// @Router /photo [GET]
func (handler PhotoHandler) ListPhoto(ctx *fiber.Ctx) error{
	res,errCode := handler.usecase.FindAll()

	if errCode == "200"{
		model.SuccessResponse(ctx,"SUCCESS GET PHOTO",res)
	}

	if errCode == "404"{
		model.NotFoundResponse(ctx,"PHOTO NOT FOUND",res)
	}

	if errCode == "500" {
		model.InternalServerErrorResponse(ctx,"INTERNAL SERVER ERROR",nil)
	}
	return nil
}


// Registeruser List Photo 
// @Summary List Photo
// @Description List Photo
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=model.PhotoResponse}
// @Failure 404 {string} model.WebResponse{code=404}
// @Failure 500 {string} model.WebResponse{code=500}
// @Router /photo/id [GET]
func (handler PhotoHandler) GetPhoto(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)

	can:=helper.Can(email,"photo.list")
	log.Print("can",can)
	
	if !can {
		model.ForbiddenResponse(ctx,"FORBIDDEN",nil)
		return nil
	}else{
		res,errCode := handler.usecase.GetPhotoById(id)

		if errCode == "200"{
			model.SuccessResponse(ctx,"SUCCESS GET PHOTO",res)
		}

		if errCode == "404"{
			model.NotFoundResponse(ctx,"PHOTO NOT FOUND",res)
		}

		if errCode == "500" {
			model.InternalServerErrorResponse(ctx,"INTERNAL SERVER ERROR",nil)
		}
	}
	return nil
}

// Registeruser Edit Photo 
// @Summary Edit Photo
// @Description Edit Photo
// @Tags photo
// @Accept json
// @Produce json
// @Param photo body model.UpdatePhotoRequest true "Update photo"
// @Success 200 {object} model.WebResponse{data=model.PhotoResponse}
// @Failure 404 {string} model.WebResponse{code=404}
// @Failure 500 {string} model.WebResponse{code=500}
// @Router /photo/id [PUT]
func (handler PhotoHandler) UpdatePhoto(ctx *fiber.Ctx) error {
	var request model.CreatePhotoRequest
	ctx.BodyParser(&request)
	id := ctx.Params("id")
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	request.Email = email
	
	result,errCode := handler.usecase.EditPhoto(id,model.UpdatePhotoRequest(request))
	log.Println(errCode)
	model.GetResponse(ctx,errCode,"",result)
	return nil
}

// DeletePhoto function removes a photo by ID
// @Summary Remove book by ID
// @Description Remove book by ID
// @Tags photo
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} model.WebResponse{}
// @Failure 404 {object} model.WebResponse{}
// @Failure 503 {object} model.WebResponse{}
// @Router /photo/{id} [delete]
func (handler PhotoHandler) DeletePhoto(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	errCode := handler.usecase.DeletePhoto(id)
	if errCode == "200"{
		model.SuccessResponse(ctx,"SUCCESS DELETE PHOTO",nil)
		return nil
	}

	if errCode == "404"{
		model.NotFoundResponse(ctx,"PHOTO NOT FOUND",nil)
		return nil
	}

	if errCode == "500" {
		model.InternalServerErrorResponse(ctx,"INTERNAL SERVER ERROR",nil)
		return nil
	}

	return nil
}