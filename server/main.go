package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "AutoTechno Server v1.0.0",
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// Health check
	app.Get("/health", healthCheck)

	// API routes
	api := app.Group("/api/v1")

	// Estaciones DI.FM
	stations := api.Group("/stations")
	stations.Get("/", getStations)
	stations.Get("/:id", getStationByID)

	// Control del reproductor
	player := api.Group("/player")
	player.Post("/play", playStation)
	player.Post("/stop", stopPlayback)
	player.Post("/pause", pausePlayback)
	player.Get("/status", getPlayerStatus)

	// Control Bluetooth
	bluetooth := api.Group("/bluetooth")
	bluetooth.Get("/devices", listBluetoothDevices)
	bluetooth.Post("/connect", connectBluetoothDevice)
	bluetooth.Post("/disconnect", disconnectBluetoothDevice)

	// Favoritos
	favorites := api.Group("/favorites")
	favorites.Get("/", getFavorites)
	favorites.Post("/", addFavorite)
	favorites.Delete("/:id", removeFavorite)

	log.Fatal(app.Listen(":3000"))
}

// Health check endpoint
func healthCheck(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
		"service": "AutoTechno Server",
	})
}

// Stations handlers
func getStations(c fiber.Ctx) error {
	stations := []fiber.Map{
		{
			"id": 1,
			"name": "Trance",
			"description": "Trance music 24/7",
		},
		{
			"id": 2,
			"name": "House",
			"description": "House music 24/7",
		},
		{
			"id": 3,
			"name": "Techno",
			"description": "Techno music 24/7",
		},
	}
	return c.JSON(fiber.Map{
		"stations": stations,
		"total": len(stations),
	})
}

func getStationByID(c fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"id": id,
		"name": "Station " + id,
		"status": "available",
	})
}

// Player handlers
func playStation(c fiber.Ctx) error {
	var req fiber.Map
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	return c.JSON(fiber.Map{
		"status": "playing",
		"station": req["station"],
		"timestamp": "now",
	})
}

func stopPlayback(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "stopped",
	})
}

func pausePlayback(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "paused",
	})
}

func getPlayerStatus(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"playing": true,
		"station": "Trance",
		"bitrate": "128kbps",
	})
}

// Bluetooth handlers
func listBluetoothDevices(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"devices": []fiber.Map{
			{"name": "Device 1", "connected": true},
			{"name": "Device 2", "connected": false},
		},
	})
}

func connectBluetoothDevice(c fiber.Ctx) error {
	var req fiber.Map
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	return c.JSON(fiber.Map{
		"status": "connected",
		"device": req["device"],
	})
}

func disconnectBluetoothDevice(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "disconnected",
	})
}

// Favorites handlers
func getFavorites(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"favorites": []fiber.Map{
			{"id": 1, "station": "Trance"},
			{"id": 2, "station": "House"},
		},
	})
}

func addFavorite(c fiber.Ctx) error {
	var req fiber.Map
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "added",
		"favorite": req,
	})
}

func removeFavorite(c fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"status": "removed",
		"id": id,
	})
}
