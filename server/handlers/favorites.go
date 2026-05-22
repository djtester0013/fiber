package handlers

import (
	"github.com/gofiber/fiber/v3"
)

// GetFavorites retrieves user's favorite stations
func GetFavorites(c fiber.Ctx) error {
	favorites := []fiber.Map{
		{
			"id":      1,
			"station": "Trance",
			"added_at": "2026-01-15T10:30:00Z",
		},
		{
			"id":      2,
			"station": "House",
			"added_at": "2026-02-20T14:45:00Z",
		},
	}

	return c.JSON(fiber.Map{
		"success":   true,
		"favorites": favorites,
		"total":     len(favorites),
	})
}

// AddFavorite adds a station to favorites
func AddFavorite(c fiber.Ctx) error {
	var req fiber.Map
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Station added to favorites",
		"data":    req,
	})
}

// RemoveFavorite removes a station from favorites
func RemoveFavorite(c fiber.Ctx) error {
	stationID := c.Params("id")

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Station removed from favorites",
		"id":      stationID,
	})
}
