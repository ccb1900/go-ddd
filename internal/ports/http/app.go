package http

import (
	// "context"
	// "log"
	// "net/http"
	// "os"
	// "os/signal"
	// "syscall"
	// "time"

	// "context"

	"context"
	"goddd/pkg/config"
	"log/slog"

	"github.com/gin-gonic/gin"
	// "github.com/spf13/viper"
)

type IRegisterRoute interface {
	RegisterRoutes(*gin.RouterGroup)
}


type App struct {
	Router *gin.Engine
	handlers []IRegisterRoute
	logger *slog.Logger
	config config.AppConfig
}

func NewApp(
	logger *slog.Logger,
	config config.AppConfig,
	handlers ...IRegisterRoute,
) *App {
	g := gin.New()
	
	for i := 0; i < len(handlers); i++ {
		handlers[i].RegisterRoutes(g.Group("/v0"))
	}
	app := &App{
		Router: g,
		logger:logger,
		config: config,
	}
	return app
}

func (app *App)Run(ctx context.Context) {
	app.Router.Run()
}
