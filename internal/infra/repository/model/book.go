package model

type BookModel struct {
	ID     int `gorm:"primarykey"`
	Name   string
	Author string
	Price  float64
}

// func (b *BookModel) ToDomain() *domain.Book {
// 	return &domain.Book{
// 		ID:     b.ID,
// 		Name:   b.Name,
// 		Author: b.Author,
// 		Price:  b.Price,
// 	}
// }
