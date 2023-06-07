package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;" column:"id"`
	UserId    uuid.UUID `gorm:"column:user_id;type:uuid;"`
	Token     string    `gorm:"column:token"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserToken) TableName() string {
	return "user_token"
}

func (entity *UserToken) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *UserToken) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
