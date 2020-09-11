package request

import (
	"net/url"
)

// Request userd to get a bat cave device setting
type GetBatCaveDeviceSettingRequest struct{}

func (r *GetBatCaveDeviceSettingRequest) Validate(id string) url.Values {
	errs := url.Values{}

	// 	regex, _ := regexp.Compile("^[a-f0-9]{12}$")
	// 	isMacAddress := regex.MatchString(deviceID)
	// 	if !isMacAddress {
	// 		errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	// 	}

	// FIXME: check for a valid uuid

	return errs
}
