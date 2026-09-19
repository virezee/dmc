-- name: CreateSensorData :exec
INSERT INTO sensor_data (temperature, humidity, created_at)
VALUES (?, ?, ?);