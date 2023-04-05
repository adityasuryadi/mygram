package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;" column:"id"`
	Message   string    `gorm:"column:message"`
	UserId    uuid.UUID `gorm:"column:user_id"`
	PhotoId   uuid.UUID `gorm:"column:photo_id"`
	User      User      `gorm:"foreignKey:UserId"`
	Photo     Photo     `gorm:"foreignKey:PhotoId"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Comment) TableName() string {
	return "comment"
}

func (entity *Comment) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Comment) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
