package application

import (
	"context"
	"goddd/internal/domain"
	"goddd/internal/ports/http/v0/dto"
)

type BookService struct {
	repo domain.IBookRepo
}

func NewBookService(r domain.IBookRepo) *BookService {
	return &BookService{
		repo: r,
	}
}

func (s *BookService) CreateBook(ctx context.Context, dto *dto.CreateBookRequest) (*dto.CreateBookResponse, error) {
	return nil, nil
}
