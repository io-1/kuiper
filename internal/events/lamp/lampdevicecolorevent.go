package lamp

type LampDeviceColorEvent struct {
	EventType string `json:"e"`
	Red       int32  `json:"r"`
	Green     int32  `json:"g"`
	Blue      int32  `json:"b"`
}

func NewLampDeviceColorEvent(red, green, blue int32) *LampDeviceColorEvent {
	return &LampDeviceColorEvent{
		EventType: "color",
		Red:       red,
		Green:     green,
		Blue:      blue,
	}
}
