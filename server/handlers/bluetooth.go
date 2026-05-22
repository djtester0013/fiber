package handlers

import (
	"github.com/gofiber/fiber/v3"
)

// ListBluetoothDevices lists all available Bluetooth devices
func ListBluetoothDevices(c fiber.Ctx) error {
	devices := []fiber.Map{
		{
			"id":        "device_1",
			"name":      "Samsung Phone",
			"connected": true,
		},
		{
			"id":        "device_2",
			"name":      "iPhone 14",
			"connected": false,
		},
		{
			"id":        "device_3",
			"name":      "Bluetooth Speaker",
			"connected": true,
		},
	}

	return c.JSON(fiber.Map{
		"success": true,
		"devices": devices,
		"total":   len(devices),
	})
}

// ConnectBluetoothDevice connects to a Bluetooth device
func ConnectBluetoothDevice(c fiber.Ctx) error {
	var req fiber.Map
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"status":  "connected",
		"device":  req["device_id"],
	})
}

// DisconnectBluetoothDevice disconnects from a Bluetooth device
func DisconnectBluetoothDevice(c fiber.Ctx) error {
	var req fiber.Map
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"status":  "disconnected",
		"device":  req["device_id"],
	})
}
