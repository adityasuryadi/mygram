package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Socialmedia struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;" column:"id"`
	Name   	  string    `gorm:"column:name"`
	UserId    uuid.UUID `gorm:"column:user_id"`
	SocialMediaUrl   string `gorm:"column:social_media_url"`
	User      User      `gorm:"foreignKey:UserId"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Socialmedia) TableName() string {
	return "socialmedia"
}

func (entity *Socialmedia) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Socialmedia) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}