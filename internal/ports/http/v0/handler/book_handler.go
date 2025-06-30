package handler

import (
	"goddd/internal/application"
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

func (h *BookHandler) Create(ctx *gin.Context) {}
func (h *BookHandler) Get(ctx *gin.Context)    {}
func (h *BookHandler) RegisterRoutes(g *gin.RouterGroup) {
	bookGroup := g.Group("/book")
	bookGroup.POST("/", h.Create)
	bookGroup.GET("/", h.Get)
}
