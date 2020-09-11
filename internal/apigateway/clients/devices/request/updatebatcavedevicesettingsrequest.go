package request

import (
	"net/url"
)

// The request used to update a Bat Cave device setting
type UpdateBatCaveDeviceSettingRequest struct {

	// The deep sleep delay of the device
	DeepSleepDelay uint32 `json:"deepSleepDelay" binding:"required"`
}

func (r *UpdateBatCaveDeviceSettingRequest) Validate(id string) url.Values {
	errs := url.Values{}

	// FIXME check for valid uuid

	// regex, _ := regexp.Compile("^[a-f0-9]{12}$")
	// isMacAddress := regex.MatchString(r.Mac)
	// if !isMacAddress {
	// 	errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	// }

	if r.DeepSleepDelay < 1 {
		errs.Add("deepSleepDelay", "The deepSleepDelay field should be a positive non-zero value!")
	}

	return errs
}
