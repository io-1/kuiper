package response

type GetAttachResponse struct {
	ID          string `json:"id"`
	ConditionID string `json:"conditionID"`
	EventID     string `json:"eventID"`
	EventType   string `json:"eventType"`
}
