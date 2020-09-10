package request

import (
	"net/url"
	"regexp"
)

// Request userd to get a bat cave device setting
// swagger:parameters getBatCaveDeviceSetting
type GetBatCaveDeviceSettingRequest struct {

	// The deviceID of the user
	//
	// in: query
	DeviceID string `json:"deviceID"`
}

func (r *GetBatCaveDeviceSettingRequest) Validate() url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("^[a-f0-9]{12}$")
	isMacAddress := regex.MatchString(r.DeviceID)
	if !isMacAddress {
		errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	}

	return errs
}
