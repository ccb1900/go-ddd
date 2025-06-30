package domain

import "errors"

var (
	ErrBookNotFound      = errors.New("book not found")
	ErrBookAlreadyExists = errors.New("book already exists")
	ErrInvalidBookState  = errors.New("invalid book state")
)

type Book struct {
	ID     int64
	Title  string
	Author string
}
type ConditionDto struct {
	Title string
	State int
}

type ISpecification[T any] interface {
	IsSatisfiedBy(T) bool
}

// 明确参数可以是基本类型和领域对象
type IBookRepo interface {
	FindByID(id int) (Book, error)
	FindAll(query ConditionDto, page, pageSize int) ([]Book, error)
	Save(book Book) error
	Delete(book Book) error
	DeleteByID(id int) error
}
