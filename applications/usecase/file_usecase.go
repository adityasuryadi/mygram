package usecase

import (
	"crypto/sha1"
	"fmt"
	"mime/multipart"
	"mygram/domains"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewFileUsecase() domains.FileUsecase {
	return &FileUsecaseImpl{}
}

type FileUsecaseImpl struct {
}

// UploadFile implements domains.FileUsecase
func (usecase *FileUsecaseImpl) UploadFile(file *multipart.FileHeader) (string, error) {
	var sha = sha1.New()
	var ctx *fiber.Ctx

	extension := path.Ext(file.Filename)
	fileName := strings.TrimSuffix(file.Filename, extension)
	sha.Write([]byte(fileName))
	encrypted := sha.Sum(nil)
	encryptedFileName := fmt.Sprintf("%x", encrypted)

	saveFile := "./public/images/" + encryptedFileName + extension
	err := ctx.SaveFile(file, saveFile)
	if err != nil {
		fmt.Println(err)
	}
	return saveFile, nil
}
