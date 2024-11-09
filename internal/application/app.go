package application

import (
	"context"
	"fmt"

	"github.com/lelikptz/golang-service-boilerplate/internal/infrastructure/logger"
)

type App struct {
	config *Config
}

func NewApp(config *Config) *App {
	return &App{config}
}

func (a *App) Run(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("Application `%s` starting!", a.config.Service.Name))
	defer logger.Info(ctx, fmt.Sprintf("Application `%s` finish!", a.config.Service.Name))

	return nil
}
