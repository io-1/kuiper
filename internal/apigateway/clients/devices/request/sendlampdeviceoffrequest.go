package request

import (
	"net/url"
	"regexp"
)

type SendLampDeviceOffRequest struct {
}

func (r SendLampDeviceOffRequest) Validate(mac string) url.Values {
	errs := url.Values{}

	macRegex, _ := regexp.Compile("^([0-9a-f]{2}){5}([0-9a-f]{2})$")
	validMac := macRegex.MatchString(mac)
	if !validMac {
		errs.Add("mac", "The mac field needs to be a valid!")
	}

	return errs
}
