package mqtt

import "github.com/eclipse/paho.mqtt.golang"

type MQTT struct {
	client mqtt.Client
}

func NewMQTT(broker, clientID string) (*MQTT, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &MQTT{client}, nil
}

func (m *MQTT) Connected() bool {
	return m.client.IsConnected()
}

func (m *MQTT) Publish(topic string, payload []byte) error {
	token := m.client.Publish(topic, 0, false, payload)
	token.Wait()
	if err := token.Error(); err != nil {
		return err
	}
	return nil
}

func (m *MQTT) Close() {
	m.client.Disconnect(250)
}
