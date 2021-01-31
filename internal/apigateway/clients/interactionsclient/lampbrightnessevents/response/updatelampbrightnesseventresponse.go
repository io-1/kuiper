package response

type UpdateLampBrightnessEventResponse struct {
	ID         string `json:"id"`
	Mac        string `json:"mac"`
	Brightness int32  `json:"brightness"`
}
