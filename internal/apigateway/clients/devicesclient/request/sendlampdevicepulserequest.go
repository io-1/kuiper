package request

import (
	"net/url"
	"regexp"
)

type SendLampDevicePulseRequest struct {
	Red   *int32 `json:"red" binding:"required"`
	Green *int32 `json:"green" binding:"required"`
	Blue  *int32 `json:"blue" binding:"required"`
}

func (r SendLampDevicePulseRequest) Validate(mac string) url.Values {
	errs := url.Values{}

	macRegex, _ := regexp.Compile("^([0-9a-f]{2}){5}([0-9a-f]{2})$")
	validMac := macRegex.MatchString(mac)
	if !validMac {
		errs.Add("mac", "The mac field needs to be a valid!")
	}

	if *r.Red < 0 {
		errs.Add("red", "The red field needs to be equal or greater then 0!")
	}

	if *r.Red > 256 {
		errs.Add("red", "The red field needs to be equal or less then 255!")
	}

	if *r.Green < 0 {
		errs.Add("green", "The green field needs to be equal or greater then 0!")
	}

	if *r.Green > 256 {
		errs.Add("green", "The green field needs to be equal or less then 255!")
	}

	if *r.Blue < 0 {
		errs.Add("blue", "The blue field needs to be equal or greater then 0!")
	}

	if *r.Blue > 256 {
		errs.Add("blue", "The blue field needs to be equal or less then 255!")
	}

	return errs
}
