package main

import (
	"avito/config"
	"avito/internal/handlers"
	"avito/internal/repository"
	"avito/internal/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	routers    *fiber.App
	repository *repository.Repository
	service    *service.Service
	handlers   *handlers.Handlers
}

func main() {
	fmt.Println("Start")
	//logger := slog.Logger{} //засунуть в контекст
	app := &App{}
	app.routers = fiber.New()
	cfg := config.Config_load()
	app.repository = repository.New(cfg)
	go func() {
		err := app.repository.CacheRecovery()
		if err != nil {
			fmt.Println("Fail in cache recovery", err)
		}
	}()
	app.service = service.New(app.repository)
	app.handlers = handlers.New(app.service)
	app.routers.Post("/", app.handlers.Post)
	app.routers.Get("/:Link", app.handlers.Get)
	err := app.routers.Listen(":3000")
	if err != nil {
		return
	}
}
