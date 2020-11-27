package response

type PatchLampEventResponse struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
	Color     string `json:"color"`
}
