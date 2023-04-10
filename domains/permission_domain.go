package domains

import entities "mygram/domains/entity"

type PermissionRepository interface {
	Insert(entites *entities.Permission) error
}

type PermissionUsecase interface {
}