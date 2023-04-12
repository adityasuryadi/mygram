package entities

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id        int `gorm:"primaryKey;type:int;autoIncrement;" column:"id"`
	Name  string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Permissions []Permission `gorm:"many2many:permission_role"`
}

func (Role) TableName() string {
	return "role"
}

func (entity *Role) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Role) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}