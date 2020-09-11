package response

// CreateBatCaveSettingResponse is a response when successfully creating a Bat Cave device setting.
type CreateBatCaveDeviceSettingResponse struct {

	// The ID of the device.
	ID string `json:"id"`

	// The device ID of the device.
	Mac string `json:"mac"`

	// The deep sleep delay of the device.
	DeepSleepDelay uint32 `json:"deepSleepDelay"`
}
