package response

type PatchAttachResponse struct {
	ID          string `json:"id"`
	ConditionID string `json:"conditionID"`
	EventID     string `json:"eventID"`
	EventType   string `json:"eventType"`
}
