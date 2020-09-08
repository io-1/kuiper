package response

// CreateBatCaveSettingResponse is a response when successfully creating a Bat Cave device setting.
// swagger:response CreateBatCaveDeviceSettingResponse
type CreateBatCaveDeviceSettingResponse struct {

	// The device ID of the device.
	DeviceID string `json:"deviceID"`

	// The deep sleep delay of the device.
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
