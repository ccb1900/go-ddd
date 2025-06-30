//go:build wireinject
// +build wireinject

package di

//go:generate go tool wire

// go:build wireinject

// internal/infra/di/wire.go

import (
	"goddd/internal/application"
	"goddd/internal/infra/database"
	"goddd/internal/infra/logger"
	"goddd/internal/ports/http"
	"goddd/internal/ports/http/v0/handler"
	"goddd/internal/infra/repository"
	"goddd/pkg/config"

	"github.com/google/wire"
)

func InitializeApp() *http.App {
	wire.Build(
		config.ProvideConfig,
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
