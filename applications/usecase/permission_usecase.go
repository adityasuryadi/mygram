package usecase

import (
	"log"
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
)

func NewPermissionUsecase(permissionRepo domains.PermissionRepository) domains.PermissionUsecase {
	return &PermissionUsecaseImpl{
		permissionRepo: permissionRepo,
	}
}

type PermissionUsecaseImpl struct {
	permissionRepo domains.PermissionRepository
}

func (PermissionUsecase *PermissionUsecaseImpl) CreatePermission(request *model.CreatePermissionRequest)(string,interface{},*model.PermissionResponse) {
	responseCode := make(chan string,1)
	permission := &entities.Permission{
		Name: request.Name,
	}
	err := PermissionUsecase.permissionRepo.Insert(permission)
	log.Print(err)
	if err != nil {
		responseCode <- "500"
		return <-responseCode,nil,nil
	}

	responseCode <- "200"
	return <-responseCode,nil,nil
}