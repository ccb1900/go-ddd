package main

import (
	"context"
	"fmt"
	"goddd/internal/infra/di"
	"goddd/pkg/config"
)

func main() {
	config.Init("config")
	fmt.Println("hello world")
	app := di.InitializeApp()

	app.Run(context.TODO())
}
