package request

import (
	"net/url"
	"regexp"
)

type PatchAttachRequest struct {
	ConditionID string `json:"conditionID"`
	EventID     string `json:"eventID"`
	EventType   string `json:"eventType"`
}

func (r PatchAttachRequest) Validate(id string) url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	validID := regex.MatchString(id)
	if !validID {
		errs.Add("id", "The id field needs to be a valid!")
	}

	validConditionID := regex.MatchString(r.ConditionID)
	if !validConditionID {
		errs.Add("conditionID", "The conditionID field needs to be a valid!")
	}

	validEventID := regex.MatchString(r.EventID)
	if !validEventID {
		errs.Add("id", "The id field needs to be a valid!")
	}

	if len(r.EventType) > 50 {
		errs.Add("eventType", "The eventType field is longer then 50 characters!")
	}

	return errs
}
