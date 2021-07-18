package sensors

type HCSR501Measurement struct {
	Mac   string `json:"m"`
	State int    `json:"p"`
}
