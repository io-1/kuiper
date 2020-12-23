package response

type PatchAttachResponse struct {
	ID            string `json:"id"`
	InteractionID string `json:"interactionID"`
	ConditionID   string `json:"conditionID"`
	EventID       string `json:"eventID"`
}
