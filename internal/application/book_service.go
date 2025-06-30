package application

import (
	"context"
	"goddd/internal/domain"
)

type BookService struct {
	repo domain.IBookRepo
}

func NewBookService(r domain.IBookRepo) *BookService {
	return &BookService{
		repo: r,
	}
}

func (s *BookService) CreateBook(ctx context.Context, dto *AddBookDtoRequest) (*AddBookDtoResponse, error) {
	s.repo.Save(domain.Book{})
	return nil, nil
}
