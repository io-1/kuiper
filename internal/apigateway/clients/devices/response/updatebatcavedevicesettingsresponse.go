package response

// UpdateBatCaveDviceSettingResonse is the response when successfully updated a Bat Cave device setting.
// swagger:response UpdateBatCaveDeviceSettingResponse
type UpdateBatCaveDeviceSettingResponse struct {

	// The device ID of the device.
	DeviceID string `json:"deviceID"`

	// The deep sleep delay of the device.
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
