package request

import (
	"net/url"
	"regexp"
)

type GetInteractionRequest struct{}

func (r GetInteractionRequest) Validate(id string) url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	validID := regex.MatchString(id)
	if !validID {
		errs.Add("id", "The id field needs to be a valid!")
	}

	if len(id) > 36 {
		errs.Add("id", "The id field is longer then 36 characters long!")
	}

	return errs
}
