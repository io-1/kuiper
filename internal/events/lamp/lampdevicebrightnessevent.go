package lamp

type LampDeviceBrightnessEvent struct {
	EventType  string `json:"e"`
	Brightness int32  `json:"br"`
}
