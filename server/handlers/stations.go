package handlers

import (
	"github.com/gofiber/fiber/v3"
)

// GetStations retrieves all DI.FM stations
func GetStations(c fiber.Ctx) error {
	stations := []fiber.Map{
		{
			"id":          1,
			"name":        "Trance",
			"description": "24/7 Trance Music",
			"url":         "https://listen.di.fm/public1/trance.pls",
		},
		{
			"id":          2,
			"name":        "House",
			"description": "24/7 House Music",
			"url":         "https://listen.di.fm/public1/house.pls",
		},
		{
			"id":          3,
			"name":        "Techno",
			"description": "24/7 Techno Music",
			"url":         "https://listen.di.fm/public1/techno.pls",
		},
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    stations,
		"total":   len(stations),
	})
}

// GetStationByID retrieves a specific station by ID
func GetStationByID(c fiber.Ctx) error {
	stationID := c.Params("id")

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":   stationID,
			"name": "Station " + stationID,
		},
	})
}
