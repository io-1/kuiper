package request

import (
	"net/url"
	"regexp"
)

type CreateAttachRequest struct {
	ConditionID string `json:"conditionID" binding:"required"`
	EventID     string `json:"eventID" binding:"required"`
}

func (r CreateAttachRequest) Validate() url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	validConditionID := regex.MatchString(r.ConditionID)
	if !validConditionID {
		errs.Add("conditionID", "The conditionID field needs to be a valid!")
	}

	validEventID := regex.MatchString(r.EventID)
	if !validEventID {
		errs.Add("id", "The id field needs to be a valid!")
	}

	return errs
}
