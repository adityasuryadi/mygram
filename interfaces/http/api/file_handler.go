package handler

import (
	"fmt"
	"mygram/domains"
	"mygram/domains/model"
	"mygram/infrastructures/validation"
	"path"

	"github.com/gofiber/fiber/v2"
)

func NewFileHandler(usecase domains.FileUsecase, validate validation.Validation) FileHandler {
	return FileHandler{
		usecase:  usecase,
		validate: validate,
	}
}

type FileHandler struct {
	usecase  domains.FileUsecase
	validate validation.Validation
}

func (handler FileHandler) Route(app *fiber.App) {
	file := app.Group("file")
	file.Post("/", handler.Upload)
}

func (handler FileHandler) Upload(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	file := form.File["gambar"][0]
	fileName := file.Filename
	extension := path.Ext(file.Filename)

	if err != nil {
		fmt.Println(err)
	}

	request := &model.CreateFileRequest{
		Name: ctx.FormValue("name"),
		Gambar: model.Image{
			Filename:    fileName,
			ContentType: "png",
			Bytes:       int32(file.Size),
			Ext:         extension,
		},
	}

	ctx.BodyParser(request)
	validationErr := handler.validate.ValidateRequest(request)

	model.GetResponse(ctx, "200", "nil", validationErr)
	return nil
}
