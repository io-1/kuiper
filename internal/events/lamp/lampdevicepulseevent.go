package lamp

type LampDevicePulseEvent struct {
	EventType string `json:"e"`
	Red       int32  `json:"r"`
	Green     int32  `json:"g"`
	Blue      int32  `json:"b"`
}

func NewLampDevicePulseEvent(red, green, blue int32) *LampDevicePulseEvent {
	return &LampDevicePulseEvent{
		EventType: "pulse",
		Red:       red,
		Green:     green,
		Blue:      blue,
	}
}
