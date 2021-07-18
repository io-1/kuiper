package response

type GetKeypadConditionToLampEventResponse struct {
	ID            string `json:"id"`
	InteractionID string `json:"interactionID"`
	ConditionID   string `json:"conditionID"`
	EventID       string `json:"eventID"`
}
