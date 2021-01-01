package response

type UpdateLampEventResponse struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
	Red       int32  `json:"red"`
	Green     int32  `json:"green"`
	Blue      int32  `json:"blue"`
}
