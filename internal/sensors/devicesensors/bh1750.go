package sensors

type BH1750Measurement struct {
	Mac       string  `json:"m"`
	Intensity float64 `json:"i"`
}
