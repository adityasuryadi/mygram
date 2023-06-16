package domains

import (
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

type UserTokenRepository interface {
	InsertTokenWithTx(tx *gorm.DB, user *entities.User, token string) error
	RemoveToken(userId string) error
}
