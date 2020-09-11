package response

type BatCaveDeviceSettingResponse struct {
	DeepSleepDelay uint32 `json:"s"`
}

func GetBatCaveDeviceSettingDefault() BatCaveDeviceSettingResponse {
	return BatCaveDeviceSettingResponse{
		DeepSleepDelay: 15,
	}
}
