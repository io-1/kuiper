package response

type LampDeviceBrightnessResponse struct {
	EventType  string `json:"e"`
	Brightness int32  `json:"b"`
}
