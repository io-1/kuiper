package lamp

type LampDeviceMeteorEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceMeteorEvent() *LampDeviceMeteorEvent {
	return &LampDeviceMeteorEvent{
		EventType: "meteor",
	}
}
