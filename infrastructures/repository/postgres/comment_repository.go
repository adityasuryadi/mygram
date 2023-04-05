package repository

import (
	"mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewCommentRepository(database *gorm.DB) domains.CommentRepository{
	return &CommentRepositoryImpl{
		db:database,
	}
}

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func (repository *CommentRepositoryImpl) StoreComment(comment entities.Comment) error{
	err := repository.db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}