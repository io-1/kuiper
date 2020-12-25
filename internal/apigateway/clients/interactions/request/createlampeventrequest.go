package request

import (
	"net/url"
	"regexp"
)

type CreateLampEventRequest struct {
	Mac       string `json:"mac" binding:"required"`
	EventType string `json:"eventType" binding:"required"`
	Red       *int32 `json:"red" binding:"required"`
	Green     *int32 `json:"green" binding:"required"`
	Blue      *int32 `json:"blue" binding:"required"`
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

	if *r.Red > 256 {
		errs.Add("red", "The red field needs to be less then 255!")
	}

	if *r.Green > 256 {
		errs.Add("green", "The green field needs to be less then 255!")
	}

	if *r.Blue > 256 {
		errs.Add("blue", "The blue field needs to be less then 255!")
	}

	return errs
}
