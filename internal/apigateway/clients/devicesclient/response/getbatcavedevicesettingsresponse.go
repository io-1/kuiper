package response

// GetBatCaveDeviceSettingResponse is a response when successfully getting a Bat Cave device setting.
type GetBatCaveDeviceSettingResponse struct {

	// The ID of the device.
	ID string `json:"id"`

	// The mac of the device.
	Mac string `json:"mac"`

	// The deep sleep delay of the device.
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
