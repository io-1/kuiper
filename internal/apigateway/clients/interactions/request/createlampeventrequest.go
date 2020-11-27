package request

import (
	"net/url"
	"regexp"
)

type CreateLampEventRequest struct {
	Mac       string `json:"mac" binding:"required"`
	EventType string `json:"eventType" binding:"required"`
	Color     string `json:"color" binding:"required"`
}

func (r CreateLampEventRequest) Validate() url.Values {
	errs := url.Values{}

	macRegex, _ := regexp.Compile("^([0-9a-f]{2}){5}([0-9a-f]{2})$")
	validMac := macRegex.MatchString(r.Mac)
	if !validMac {
		errs.Add("mac", "The mac field needs to be a valid!")
	}

	if len(r.EventType) > 50 {
		errs.Add("eventType", "The eventType field needs to be less then 50 characters!")
	}

	if len(r.Color) > 10 {
		errs.Add("color", "The color field needs to be less then 10 characters!")
	}

	return errs
}
