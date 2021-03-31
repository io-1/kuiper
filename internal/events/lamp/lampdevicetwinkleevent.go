package lamp

type LampDeviceTwinkleEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceTwinkleEvent() *LampDeviceOnEvent {
	return &LampDeviceOnEvent{
		EventType: "twinkle",
	}
}
