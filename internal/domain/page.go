package domain

type PageRequest struct {
	Page     int
	PageSize int
}

type PageResult[T any] struct {
	Items []T
	Total int
}
