package response

type UpdateLampEventResponse struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
	Red       int32  `json:"red" binding:"required"`
	Green     int32  `json:"green" binding:"required"`
	Blue      int32  `json:"blue" binding:"required"`
}
