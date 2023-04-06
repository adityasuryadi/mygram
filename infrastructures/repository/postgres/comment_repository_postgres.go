package repository

import (
	"errors"
	"mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewCommentRepository(database *gorm.DB) domains.CommentRepository {
	return &CommentRepositoryImpl{
		db: database,
	}
}

type CommentRepositoryImpl struct {
	db *gorm.DB
}

// DestroyComment implements domains.CommentRepository
func (repository *CommentRepositoryImpl) DestroyComment(id string) error {
	err := repository.db.Where("id = ?",id).Delete(&entities.Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateComment implements domains.CommentRepository
func (repository *CommentRepositoryImpl) UpdateComment(id string, comment entities.Comment) error {
	err := repository.db.Save(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

// FindCommentById implements domains.CommentRepository
func (repository *CommentRepositoryImpl) FindCommentById(id string) (*entities.Comment, error) {
	var comment entities.Comment
	err := repository.db.Where("id = ?", id).First(&comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &comment, nil
}

// GetComment implements domains.CommentRepository
func (repository *CommentRepositoryImpl) GetComment() (comments []*entities.Comment, err error) {
	err = repository.db.Find(&comments).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return comments, nil
}

func (repository *CommentRepositoryImpl) StoreComment(comment entities.Comment) error {
	err := repository.db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}
