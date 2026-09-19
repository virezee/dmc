package inbound

import (
	"context"

	"dmc/internal/domain/device"
)

type DeviceUseCase interface {
	SendCommand(ctx context.Context, device device.Device) error
}
