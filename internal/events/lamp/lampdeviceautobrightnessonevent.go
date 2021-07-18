package lamp

type LampDeviceAutoBrightnessOnEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceAutoBrightnessOnEvent() *LampDeviceAutoBrightnessToggleEvent {
	return &LampDeviceAutoBrightnessToggleEvent{
		EventType: "auto-brightness-on",
	}
}
