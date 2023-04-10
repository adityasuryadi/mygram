package entities

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	Id        int `gorm:"primaryKey;type:int;autoIncrement;" column:"id"`
	Name  		string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Permission) TableName() string {
	return "Permission"
}

func (entity *Permission) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Permission) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}