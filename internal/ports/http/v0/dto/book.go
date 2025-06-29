package dto

import "goddd/internal/domain"

type GetBookRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetBookResponse struct {
	Book domain.Book `json:"book"`
}


type CreateBookRequest struct {
    Title  string `json:"title" binding:"required"`  // 请求字段校验
    ISBN   string `json:"isbn" binding:"required"`
    Author string `json:"author"`
}

type CreateBookResponse struct {
	ID string
}

// 可添加转换方法（非必须）
func (d *CreateBookRequest) ToDomainEntity() *domain.Book {
    return &domain.Book{
        Title:  d.Title,
        // ISBN:   domain.ISBN(d.ISBN), // 转换为值对象
        Author: d.Author,
    }
}