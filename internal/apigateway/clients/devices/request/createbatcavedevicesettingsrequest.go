package request

import (
	"net/url"
	"regexp"
)

// The request used to create a Bat Cave device setting
type CreateBatCaveDeviceSettingRequest struct {

	// The mac of the device
	Mac string `json:"mac" binding:"required"`

	// The deep sleep delay of the device
	DeepSleepDelay uint32 `json:"deepSleepDelay" binding:"required"`
}

func (r *CreateBatCaveDeviceSettingRequest) Validate() url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("^[a-f0-9]{12}$")
	isMacAddress := regex.MatchString(r.Mac)
	if !isMacAddress {
		errs.Add("mac", "The mac field needs to be a valid mac!")
	}

	if r.DeepSleepDelay < 1 {
		errs.Add("deepSleepDelay", "The deepSleepDelay field should be a positive non-zero value!")
	}

	return errs
}
