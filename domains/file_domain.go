package domains

import (
	"mime/multipart"
)

type FileUsecase interface {
	UploadFile(file *multipart.FileHeader) (string, error)
}
