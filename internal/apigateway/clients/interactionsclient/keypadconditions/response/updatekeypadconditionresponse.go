package response

type UpdateKeypadConditionResponse struct {
	ID       string `json:"id"`
	Mac      string `json:"mac"`
	ButtonID int32  `json:"buttonID"`
}
