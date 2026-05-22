package handlers

import (
	"github.com/gofiber/fiber/v3"
)

// PlayStation plays a specific station
func PlayStation(c fiber.Ctx) error {
	var req fiber.Map
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"status":  "playing",
		"station": req["station"],
	})
}

// StopPlayback stops the current playback
func StopPlayback(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"status":  "stopped",
	})
}

// PausePlayback pauses the current playback
func PausePlayback(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"status":  "paused",
	})
}

// GetPlayerStatus returns the current player status
func GetPlayerStatus(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"status": fiber.Map{
			"playing":   true,
			"station":   "Trance",
			"bitrate":   "128kbps",
			"timestamp": "2026-05-22T10:00:00Z",
		},
	})
}
