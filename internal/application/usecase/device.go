package usecase

import (
	"context"

	"dmc/internal/application/port/outbound"
	"dmc/internal/domain/device"
)

type deviceUseCase struct {
	mqtt outbound.MQTT
}

func NewDeviceUseCase(mqtt outbound.MQTT) *deviceUseCase {
	return &deviceUseCase{mqtt}
}

func (u *deviceUseCase) SendCommand(ctx context.Context, device device.Device) error {
	return u.mqtt.PublishCommand(device)
}
