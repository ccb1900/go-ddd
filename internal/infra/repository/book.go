package repository

import (
	"context"
	"errors"
	"goddd/internal/domain"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

// new
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

// Get
func (r *BookRepository) Get(ctx context.Context,id int) (*domain.Book, error) {
	var book domain.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// Update
func (r *BookRepository) Update(ctx context.Context,book *domain.Book) error {
	// 应该判断影响行数
	res := r.db.Save(book)
	if book.ID == 0 {
		return errors.New("book id is required")
	}
	if res.RowsAffected == 0 {
		return errors.New("book not found")
	}

	return res.Error
}
