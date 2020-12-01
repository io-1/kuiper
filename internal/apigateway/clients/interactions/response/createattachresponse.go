package response

type CreateAttachResponse struct {
	ID          string `json:"id"`
	ConditionID string `json:"conditionID"`
	EventID     string `json:"eventID"`
}
