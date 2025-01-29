package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var startTime = time.Now()

func HealthCheck(ctx *fiber.Ctx) error {
	uptime := time.Since(startTime)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Server is online",
		"uptime": fmt.Sprintf("%02dh %02dm %02ds", int(uptime.Hours()), int(uptime.Minutes()), int(uptime.Seconds())),
	})
}