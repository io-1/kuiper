package response

type PatchKeypadConditionResponse struct {
	ID       string `json:"id"`
	Mac      string `json:"mac"`
	ButtonID int32  `json:"buttonID"`
}
