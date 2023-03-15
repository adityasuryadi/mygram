package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id            uuid.UUID       `gorm:"primaryKey;type:uuid;" column:"id"`
	UserName      string          `gorm:"column:username"`
	Email         string          `gorm:"column:email"`
	Password      string          `gorm:"column:password"`
	Age 		  int			  `gorm:"column:age"`
	CreatedAt     time.Time       `gorm:"column:created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
