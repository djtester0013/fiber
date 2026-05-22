# AutoTechno Server - Fiber Backend

REST API server built with [Fiber](https://gofiber.io) for the AutoTechno application. This server provides endpoints for controlling DI.FM stations, managing player state, and handling Bluetooth device connections.

## 🚀 Features

- **Stations Management** - Get DI.FM radio stations
- **Player Control** - Play, pause, stop stations
- **Bluetooth Integration** - List and connect devices
- **Favorites Management** - Save/remove favorite stations
- **CORS Enabled** - Support for cross-origin requests
- **Logging Middleware** - HTTP request logging

## 📋 Prerequisites

- Go 1.25 or higher
- Git

## 🔧 Installation

```bash
cd server
go mod download
go mod tidy
```

## ▶️ Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:3000`

## 📚 API Endpoints

### Health Check
```
GET /health
```

### Stations
```
GET    /api/v1/stations          # Get all stations
GET    /api/v1/stations/:id      # Get station by ID
```

### Player Control
```
POST   /api/v1/player/play       # Start playing
POST   /api/v1/player/stop       # Stop playback
POST   /api/v1/player/pause      # Pause playback
GET    /api/v1/player/status     # Get player status
```

### Bluetooth
```
GET    /api/v1/bluetooth/devices         # List devices
POST   /api/v1/bluetooth/connect         # Connect device
POST   /api/v1/bluetooth/disconnect      # Disconnect device
```

### Favorites
```
GET    /api/v1/favorites         # Get favorites
POST   /api/v1/favorites         # Add favorite
DELETE /api/v1/favorites/:id     # Remove favorite
```

## 🧪 Testing with cURL

```bash
# Health check
curl http://localhost:3000/health

# Get all stations
curl http://localhost:3000/api/v1/stations

# Play a station
curl -X POST http://localhost:3000/api/v1/player/play \
  -H "Content-Type: application/json" \
  -d '{"station": "Trance"}'

# Get Bluetooth devices
curl http://localhost:3000/api/v1/bluetooth/devices

# Add favorite
curl -X POST http://localhost:3000/api/v1/favorites \
  -H "Content-Type: application/json" \
  -d '{"station": "House"}'
```

## 📁 Project Structure

```
server/
├── main.go              # Application entry point
├── handlers/
│   ├── stations.go      # Station endpoints
│   ├── player.go        # Player control endpoints
│   ├── bluetooth.go     # Bluetooth endpoints
│   └── favorites.go     # Favorites endpoints
├── go.mod              # Go module definition
├── .gitignore          # Git ignore rules
└── README.md           # This file
```

## 🔌 Integration with AutoTechno App

The app can communicate with this server via HTTP requests:

```kotlin
// Kotlin example for Android
val client = OkHttpClient()
val request = Request.Builder()
    .url("http://your-server:3000/api/v1/stations")
    .build()
```

## 📦 Dependencies

- [Fiber v3](https://github.com/gofiber/fiber) - Web framework
- Standard Go libraries

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

MIT License - see LICENSE file for details

## 📞 Support

For issues and questions, please open an issue on GitHub.
