package pubsub

type PubSub interface {
	NewBatCaveDeviceSettingsListener(listenerName string, mqttURL string) error
}
