package lamp

type LampDeviceAutoBrightnessOffEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceAutoBrightnessOffEvent() *LampDeviceAutoBrightnessToggleEvent {
	return &LampDeviceAutoBrightnessToggleEvent{
		EventType: "auto-brightness-off",
	}
}
