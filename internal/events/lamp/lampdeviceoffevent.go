package lamp

type LampDeviceOffEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceOffEvent() *LampDeviceOffEvent {
	return &LampDeviceOffEvent{
		EventType: "off",
	}
}
