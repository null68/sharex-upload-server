package main

import (
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/null68/sharex-upload-server/config"
	"github.com/null68/sharex-upload-server/internal/handlers"
)

func main() {

	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		return
	}

	var app (*fiber.App) = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,

		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Get("/health", handlers.HealthCheck)

	app.Listen(":" + strconv.Itoa(cfg.Port))
}