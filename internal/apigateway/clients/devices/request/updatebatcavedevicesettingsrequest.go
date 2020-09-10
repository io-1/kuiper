package request

import (
	"net/url"
)

// The request used to update a Bat Cave device setting
type UpdateBatCaveDeviceSettingRequest struct {

	// The deep sleep deley of the device
	DeepSleepDelay uint32 `json:"deepSleepDelay" binding:"required"`
}

// swagger:parameters UpdateBatCaveDeviceSettingRequest updateBatCaveDeviceSetting
type UpdateBatCaveDeviceSettingRequestWrapper struct {

	// The username of the user
	//
	// in: query
	DeviceID string `json:"deviceID"`

	// in: body
	Body UpdateBatCaveDeviceSettingRequest
}

func (r *UpdateBatCaveDeviceSettingRequest) Validate() url.Values {
	errs := url.Values{}

	// regex, _ := regexp.Compile("^[a-f0-9]{12}$")
	// isMacAddress := regex.MatchString(r.DeviceID)
	// if !isMacAddress {
	// 	errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	// }

	if r.DeepSleepDelay < 1 {
		errs.Add("deepSleepDelay", "The deepSleepDelay field should be a positive non-zero value!")
	}

	return errs
}
