package response

type PatchKeypadConditionToLampEventResponse struct {
	ID            string `json:"id"`
	InteractionID string `json:"interactionID"`
	ConditionID   string `json:"conditionID"`
	EventID       string `json:"eventID"`
}
