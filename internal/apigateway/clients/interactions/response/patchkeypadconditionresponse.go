package response

type PatchKeypadConditionResponse struct {
	ID            string `json:"id"`
	InteractionID string `json:"interactionID"`
	Mac           string `json:"mac"`
	ButtonID      int32  `json:"buttonID"`
}
