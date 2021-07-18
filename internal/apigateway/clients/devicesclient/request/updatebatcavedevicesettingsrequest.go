package request

import (
	"net/url"
	"regexp"
)

// The request used to update a Bat Cave device setting
type UpdateBatCaveDeviceSettingRequest struct {

	// The deep sleep delay of the device
	DeepSleepDelay uint32 `json:"deepSleepDelay" binding:"required"`
}

func (r *UpdateBatCaveDeviceSettingRequest) Validate(id string) url.Values {
	errs := url.Values{}

	// check for a valid uuid
	regex, _ := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	isMacAddress := regex.MatchString(id)
	if !isMacAddress {
		errs.Add("id", "The id field needs to be a valid!")
	}

	if r.DeepSleepDelay < 1 {
		errs.Add("deepSleepDelay", "The deepSleepDelay field should be a positive non-zero value!")
	}

	return errs
}
