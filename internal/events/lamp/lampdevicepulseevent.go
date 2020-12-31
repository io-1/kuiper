package lamp

type LampDevicePulseEvent struct {
	EventType string `json:"e"`
	Red       int32  `json:"r"`
	Green     int32  `json:"g"`
	Blue      int32  `json:"b"`
}
