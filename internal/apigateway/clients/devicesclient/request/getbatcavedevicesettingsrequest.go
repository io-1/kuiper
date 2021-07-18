package request

import (
	"net/url"
	"regexp"
)

// Request userd to get a bat cave device setting
type GetBatCaveDeviceSettingRequest struct{}

func (r *GetBatCaveDeviceSettingRequest) Validate(id string) url.Values {
	errs := url.Values{}

	// check for a valid uuid
	regex, _ := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	validID := regex.MatchString(id)
	if !validID {
		errs.Add("id", "The id field needs to be a valid!")
	}

	return errs
}
