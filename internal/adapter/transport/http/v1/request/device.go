package requestv1

type DeviceControl struct {
	DeviceID string `json:"device_id" validate:"required"`
	Command  string `json:"command"   validate:"required,oneof=ON OFF"`
}
