package response

type PatchLampBrightnessEventResponse struct {
	ID         string `json:"id"`
	Mac        string `json:"mac"`
	Brightness int32  `json:"brightness"`
}
