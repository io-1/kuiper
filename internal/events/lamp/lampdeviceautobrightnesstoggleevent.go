package lamp

type LampDeviceAutoBrightnessToggleEvent struct {
	EventType string `json:"e"`
}

func NewLampDeviceAutoBrightnessToggleEvent() *LampDeviceAutoBrightnessToggleEvent {
	return &LampDeviceAutoBrightnessToggleEvent{
		EventType: "auto-brightness-toggle",
	}
}
