package application

import (
	"context"
	"goddd/internal/ports/http/v0/dto"
	"goddd/internal/infra/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(r *repository.BookRepository) *BookService {
	return &BookService{
		repo: r,
	}
}

func (s *BookService) CreateBook(ctx context.Context, dto *dto.CreateBookRequest) (*dto.CreateBookResponse,error) {
    return nil,nil
}