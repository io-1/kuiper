package lamp

type LampDeviceFireEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceFireEvent() *LampDeviceOnEvent {
	return &LampDeviceOnEvent{
		EventType: "fire",
	}
}
