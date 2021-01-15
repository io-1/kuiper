package request

import (
	"net/url"
	"regexp"
)

type SendLampDeviceBrightnessRequest struct {
	Brightness *int32 `json:"brightness" binding:"required"`
}

func (r SendLampDeviceBrightnessRequest) Validate(mac string) url.Values {
	errs := url.Values{}

	macRegex, _ := regexp.Compile("^([0-9a-f]{2}){5}([0-9a-f]{2})$")
	validMac := macRegex.MatchString(mac)
	if !validMac {
		errs.Add("mac", "The mac field needs to be a valid!")
	}

	if *r.Brightness < 0 {
		errs.Add("brightness", "The brightness field needs to be equal or greater then 0!")
	}

	if *r.Brightness > 100 {
		errs.Add("brightness", "The brightness field needs to be equal or less then 100!")
	}

	return errs
}
