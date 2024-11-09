package main

import (
	"context"
	"log"
	"os"

	"github.com/lelikptz/golang-service-boilerplate/internal/application"
	"github.com/lelikptz/golang-service-boilerplate/internal/infrastructure/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var devEnvironment = "development"

func main() {
	config := parseConfig()
	l := newLogger(config)
	defer func() { _ = l.Sync() }()

	err := application.NewApp(config).Run(logger.WithLogger(context.Background(), l))
	if err != nil {
		l.Error(err)
	}
}

func parseConfig() *application.Config {
	config := application.Config{
		Service: application.Service{
			Environment: devEnvironment,
		},
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal([]byte(os.ExpandEnv(string(file))), &config); err != nil {
		log.Fatal(err)
	}

	return &config
}

func newLogger(config *application.Config) *zap.SugaredLogger {
	var l *zap.Logger
	if config.Service.Environment == devEnvironment {
		l = zap.Must(zap.NewDevelopment())
	} else {
		l = zap.Must(zap.NewProduction())
	}

	return l.Sugar()
}
