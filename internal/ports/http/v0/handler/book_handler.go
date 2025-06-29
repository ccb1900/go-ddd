package handler

import (
	"goddd/internal/application"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service *application.BookService
}

func NewBookHandler(s *application.BookService)*BookHandler {
	return &BookHandler{
		service: s,
	}
}

func (h *BookHandler) Create(ctx *gin.Context){}
func (h *BookHandler) Get(ctx *gin.Context){}
func (h *BookHandler) RegisterRoutes(g *gin.RouterGroup){
	bookGroup := g.Group("/book")
	bookGroup.POST("/",h.Create)
	bookGroup.GET("/",h.Get)
}

