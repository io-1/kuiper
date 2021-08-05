package lamp

type LampDeviceFireEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceFireEvent() *LampDeviceFireEvent {
	return &LampDeviceFireEvent{
		EventType: "fire",
	}
}
