package response

// GetBatCaveDeviceSettingResponse is a response when successfully getting a Bat Cave device setting.
type GetBatCaveDeviceSettingResponse struct {

	// The device ID of the device.
	DeviceID string `json:"deviceID"`

	// The deep sleep delay of the device.
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
