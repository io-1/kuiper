package response

type PatchLampPulseEventResponse struct {
	ID    string `json:"id"`
	Mac   string `json:"mac"`
	Red   int32  `json:"red"`
	Green int32  `json:"green"`
	Blue  int32  `json:"blue"`
}
