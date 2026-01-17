# RideQwik API

A RESTful API built with Go and Gin framework for a ride-sharing application.

## Features

- ✅ RESTful API with Gin framework
- ✅ Clean architecture with separation of concerns
- ✅ Environment-based configuration
- ✅ CORS middleware
- ✅ Error handling middleware
- ✅ Health check endpoint
- ✅ Structured logging
- 🔄 JWT authentication (ready to implement)
- 🔄 Database integration (ready to implement)

## Project Structure

```
.
├── main.go                  # Application entry point
├── go.mod                   # Go module dependencies
├── .env                     # Environment variables
├── .env.example            # Environment variables template
└── internal/               # Private application code
    ├── config/             # Configuration management
    │   └── config.go
    ├── handlers/           # HTTP request handlers
    │   ├── health.go
    │   ├── ride_handler.go
    │   ├── user_handler.go
    │   └── driver_handler.go
    ├── middleware/         # HTTP middleware
    │   ├── cors.go
    │   ├── error_handler.go
    │   └── auth.go
    ├── models/             # Data models
    │   ├── ride.go
    │   ├── user.go
    │   └── driver.go
    └── routes/             # Route definitions
        └── routes.go
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd RideQwik
```

2. Install dependencies:
```bash
go mod download
```

3. Copy the environment file:
```bash
cp .env.example .env
```

4. Update the `.env` file with your configuration

### Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Building the Application

```bash
go build -o rideqwik-api
```

## API Endpoints

### Health Check
- `GET /health` - Check API health status

### Rides
- `GET /api/v1/rides` - Get all rides
- `GET /api/v1/rides/:id` - Get ride by ID
- `POST /api/v1/rides` - Create a new ride
- `PUT /api/v1/rides/:id` - Update a ride
- `DELETE /api/v1/rides/:id` - Delete a ride

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create a new user
- `PUT /api/v1/users/:id` - Update a user
- `DELETE /api/v1/users/:id` - Delete a user

### Drivers
- `GET /api/v1/drivers` - Get all drivers
- `GET /api/v1/drivers/:id` - Get driver by ID
- `POST /api/v1/drivers` - Create a new driver
- `PUT /api/v1/drivers/:id` - Update a driver
- `DELETE /api/v1/drivers/:id` - Delete a driver

## Example API Calls

### Health Check
```bash
curl http://localhost:8080/health
```

### Create a Ride
```bash
curl -X POST http://localhost:8080/api/v1/rides \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "pickup_location": "123 Main St",
    "dropoff_location": "456 Oak Ave"
  }'
```

### Get All Rides
```bash
curl http://localhost:8080/api/v1/rides
```

## Development

### Running Tests
```bash
go test ./...
```

### Running with Hot Reload (using air)
```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

## Next Steps

1. **Database Integration**: Add PostgreSQL or MongoDB
2. **Authentication**: Implement JWT-based authentication
3. **Validation**: Add request validation
4. **Testing**: Write unit and integration tests
5. **Documentation**: Add Swagger/OpenAPI documentation
6. **Logging**: Implement structured logging with zerolog or zap
7. **Rate Limiting**: Add rate limiting middleware
8. **WebSockets**: Add real-time features for ride tracking

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [godotenv](https://github.com/joho/godotenv) - Environment variable management

## License

MIT License

