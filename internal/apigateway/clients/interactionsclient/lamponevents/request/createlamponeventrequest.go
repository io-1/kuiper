package request

import (
	"net/url"
	"regexp"
)

type CreateLampOnEventRequest struct {
	Mac string `json:"mac" binding:"required"`
}

func (r CreateLampOnEventRequest) Validate() url.Values {
	errs := url.Values{}

	macRegex, _ := regexp.Compile("^([0-9a-f]{2}){5}([0-9a-f]{2})$")
	validMac := macRegex.MatchString(r.Mac)
	if !validMac {
		errs.Add("mac", "The mac field needs to be a valid!")
	}

	return errs
}
