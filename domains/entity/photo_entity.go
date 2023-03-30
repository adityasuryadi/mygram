package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	Id        		uuid.UUID `gorm:"primaryKey;type:uuid;" column:"id"`
	Title     		string    `gorm:"column:title"`
	Caption   		string    `gorm:"column:caption"`
	PhotoUrl        string       `gorm:"column:photo_url"`
	UserId uuid.UUID `gorm:"column:user_id"`
	User User `gorm:"foreignKey:UserId"`;
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Photo) TableName() string {
	return "photo"
}

func (entity *Photo) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Photo) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}