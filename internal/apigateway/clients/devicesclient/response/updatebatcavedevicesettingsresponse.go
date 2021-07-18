package response

// UpdateBatCaveDviceSettingResonse is the response when successfully updated a Bat Cave device setting.
type UpdateBatCaveDeviceSettingResponse struct {

	// The ID of the device.
	ID string `json:"id"`

	// The deep sleep delay of the device.
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
