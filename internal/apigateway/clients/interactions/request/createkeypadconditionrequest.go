package request

import (
	"net/url"
	"regexp"
)

type CreateKeypadConditionRequest struct {
	InteractionID string `json:"interactionID" binding:"required"`
	Mac           string `json:"mac" binding:"required"`
	ButtonID      int32  `json:"buttonID" binding:"required"`
}

func (r CreateKeypadConditionRequest) Validate() url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	validID := regex.MatchString(r.InteractionID)
	if !validID {
		errs.Add("id", "The id field needs to be a valid!")
	}

	if len(r.InteractionID) > 36 {
		errs.Add("id", "The id field is longer then 36 characters long!")
	}

	return errs
}
