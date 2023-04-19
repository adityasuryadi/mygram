package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;" column:"id"`
	Name      string    `gorm:"column:name"`
	Stock     int       `gorm:"column:stock"`
	UserId    uuid.UUID `gorm:"column:user_id"`
	User      User      `gorm:"foreignKey:UserId"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

func (Product) TableName() string {
	return "product"
}

func (entity *Product) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Product) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}