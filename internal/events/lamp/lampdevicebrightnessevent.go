package lamp

type LampDeviceBrightnessEvent struct {
	EventType  string `json:"e"`
	Brightness int32  `json:"br"`
}

func NewLampDeviceBrightnessEvent(brightness int32) *LampDeviceBrightnessEvent {
	return &LampDeviceBrightnessEvent{
		EventType:  "brightness",
		Brightness: brightness,
	}
}
