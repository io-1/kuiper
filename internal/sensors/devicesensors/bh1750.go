package sensors

type BH1750Measurement struct {
	Mac       string `json:"m"`
	Intensity int32  `json:"i"`
}
