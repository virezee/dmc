package requestv1

type SensorData struct {
	Temperature float64 `json:"temperature" validate:"required"`
	Humidity    float64 `json:"humidity"    validate:"required"`
}
