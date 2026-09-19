package outbound

import "dmc/internal/domain/device"

type MQTT interface {
	IsConnected() bool
	PublishCommand(device device.Device) error
}
