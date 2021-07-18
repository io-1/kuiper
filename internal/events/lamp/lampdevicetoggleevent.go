package lamp

type LampDeviceToggleEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceToggleEvent() *LampDeviceToggleEvent {
	return &LampDeviceToggleEvent{
		EventType: "toggle",
	}
}
