package handler

import (
	"context"
	"goddd/internal/application"
	"goddd/internal/ports/http/v0/dto"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service *application.BookService
	logger  *slog.Logger
}

func NewBookHandler(s *application.BookService, logger *slog.Logger) *BookHandler {
	return &BookHandler{
		service: s,
		logger:  logger.WithGroup("book"),
	}
}

func (h *BookHandler) Create(ctx *gin.Context) {
	var req dto.CreateBookRequest
	h.service.CreateBook(context.TODO(), req.ToAddRequest())
}
func (h *BookHandler) Get(ctx *gin.Context) {}
func (h *BookHandler) RegisterRoutes(g *gin.RouterGroup) {
	bookGroup := g.Group("/book")
	bookGroup.POST("/", h.Create)
	bookGroup.GET("/", h.Get)
}
