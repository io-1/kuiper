package lamp

type LampDeviceOnEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceOnEvent() *LampDeviceOnEvent {
	return &LampDeviceOnEvent{
		EventType: "on",
	}
}
