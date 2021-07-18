package request

import (
	"net/url"
	"regexp"
)

// The request used to update a user
type PatchKeypadConditionRequest struct {
	Mac      string `json:"mac"`
	ButtonID int32  `json:"buttonID"`
}

func (r *PatchKeypadConditionRequest) Validate(id string) url.Values {
	errs := url.Values{}

	macRegex, _ := regexp.Compile("^([0-9a-f]{2}){5}([0-9a-f]{2})$")
	validMac := macRegex.MatchString(r.Mac)
	if !validMac {
		errs.Add("mac", "The mac field needs to be a valid!")
	}

	return errs
}
