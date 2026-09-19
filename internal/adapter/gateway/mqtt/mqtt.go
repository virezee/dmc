package mqttclient

import (
	"dmc/internal/domain/device"
	"dmc/internal/infrastructure/mqtt"
)

type MQTTClient struct {
	client *mqtt.MQTT
}

func NewMQTTClient(client *mqtt.MQTT) *MQTTClient {
	return &MQTTClient{client}
}

func (m *MQTTClient) PublishCommand(device device.Device) error {
	topic := "greenhouse/control/" + device.ID
	return m.client.Publish(topic, []byte(device.Command))
}

func (m *MQTTClient) IsConnected() bool {
	return m.client.Connected()
}
