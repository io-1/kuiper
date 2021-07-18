package request

import (
	"net/url"
	"regexp"
)

type CreateKeypadConditionRequest struct {
	Mac      string `json:"mac" binding:"required"`
	ButtonID *int32 `json:"buttonID" binding:"required"`
}

func (r CreateKeypadConditionRequest) Validate() url.Values {
	errs := url.Values{}

	macRegex, _ := regexp.Compile("^([0-9a-f]{2}){5}([0-9a-f]{2})$")
	validMac := macRegex.MatchString(r.Mac)
	if !validMac {
		errs.Add("mac", "The mac field needs to be a valid!")
	}

	if *r.ButtonID > 20 {
		errs.Add("buttonID", "The buttonID field needs to be less then 20!")
	}

	return errs
}
