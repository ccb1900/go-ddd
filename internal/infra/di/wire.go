// go:build wireinject
// +build wireinject

//go:generate go tool wire

// internal/infra/di/wire.go
package di

import (
	"goddd/internal/application"
	"goddd/internal/domain"
	"goddd/internal/infra/database"
	"goddd/internal/infra/logger"
	"goddd/internal/infra/repository"
	"goddd/internal/ports/http"
	"goddd/internal/ports/http/v0/handler"
	"goddd/pkg/config"

	"github.com/google/wire"
)

func InitializeApp() *http.App {
	wire.Build(
		config.ProvideConfig,
		wire.Bind(new(domain.IBookRepo), new(*repository.BookRepository)),
		database.NewDB,
		logger.NewLogger,
		repository.NewBookRepository,
		application.NewBookService,
		handler.NewBookHandler,
		http.ProvideV0Routers,
		http.NewApp,
	) // 声明依赖关系
	return nil // 返回值无实际意义
}
