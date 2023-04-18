package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"
)

type PermissionRepository interface {
	Insert(entites *entities.Permission) error
}

type PermissionUsecase interface {
	CreatePermission(request *model.CreatePermissionRequest) (string,interface{},*model.PermissionResponse)
}