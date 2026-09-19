## How to Run

### Requirements
- Docker

### Run
`docker compose up --build`

### Access
API: http://localhost:3000

### Stop
`docker compose down`

## System Architecture

The application uses a layered architecture that separates HTTP handling, business logic, persistence, and MQTT communication.

                         Client
                           |
                           v
                          API
                           |
                           v
                      HTTP Handler
                           |
                           v
                        Use Case
                           |
                 +---------+---------+
                 |                   |
                 v                   v
             Repository          MQTT Client
                 |                   |
                 v                   v
            SQLite DB           MQTT Broker
                                      |
                                      v
                              Greenhouse Device

The software architecture/architectural pattern uses a hybrid architecture combining Domain-Driven Design (DDD) principles, Hexagonal Architecture (Ports and Adapters), and Clean Architecture.

The main components are HTTP Handlers, Use Cases, Domain Models, Repositories, MQTT Client, MQTT Broker, and SQLite Database.

This separation keeps business logic independent from external infrastructure and makes the system easier to test, maintain, and extend.

### Sensor Data Flow

  Client / IoT System
          |
          | POST /sensor-data
          v
      HTTP Handler
          |
          v
       Use Case
          |
          v
   SQLite Repository
          |
          v
    SQLite Database

The /sensor-data endpoint validates the incoming payload and stores the sensor data in SQLite.

### Device Control Flow

    Client
      |
      | POST /device-control
      v
  HTTP Handler
      |
      v
   Use Case
      |
      v
  MQTT Client
      |
      | Publish
      v
  MQTT Broker
      |
      v
Greenhouse Device

## MQTT Integration

Mosquitto is used as the local MQTT broker for greenhouse device control.

The /device-control endpoint accepts a device ID and a command:

    {
      "device_id": "fan-01",
      "command": "ON"
    }

The backend publishes the command to the structured MQTT topic:

    greenhouse/control/{device_id}

Example:

    greenhouse/control/fan-01

The payload is either:

    ON

or:

    OFF

When running with Docker Compose, the API connects to the MQTT broker through the Docker Compose network.

## Error Handling and Edge Cases

The API validates incoming requests before executing business logic.

The following cases are handled:
- Invalid JSON payload
- Missing required fields
- Invalid field types
- Invalid device commands
- Database errors
- MQTT connection errors
- MQTT publish errors
- Database unavailable
- MQTT broker unavailable

For /device-control, only ON and OFF commands are accepted.

Database and MQTT failures are returned as errors instead of being silently ignored.

The /status endpoint reports the status of:
- Backend service
- Database connection
- MQTT connection

## API Endpoints

### POST /sensor-data

Example request:

    {
      "temperature": 28.5,
      "humidity": 72
    }

### POST /device-control

Example request:

    {
      "device_id": "fan-01",
      "command": "ON"
    }

### GET /status

Example response:

    {
      "data": {
        "backend": true,
        "database": true,
        "mqtt": true
      }
    }

## Design Decisions

### SQLite

SQLite was selected because the assignment allows a local database and does not require a separate database server.

### MQTT

MQTT is used for device control because its publish/subscribe model is lightweight and suitable for IoT communication.

### DDD + Hexa + Clean

HTTP handlers, use cases, database access, and MQTT communication are separated into different layers. This keeps responsibilities isolated and makes the application easier to test and maintain.

### SQLC

SQLC is used to generate type-safe database access code from SQL queries.

### Graceful Shutdown

The application closes the MQTT and database connections during shutdown to release resources cleanly.

## Database

The application uses SQLite:

    data/dmc.db

Database migrations are maintained in:

    migrations/

The database file included in the project is already initialized, so the application does not need to execute migrations every time it starts.

## Testing

Run the test suite with:

    go test ./...

The tests cover sensor data ingestion, device control, MQTT publishing, status checking, and request validation.