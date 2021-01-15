package request

import (
	"net/url"
	"regexp"
)

// The request used to update a user
type PatchInteractionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (r *PatchInteractionRequest) Validate(id string) url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	validID := regex.MatchString(id)
	if !validID {
		errs.Add("id", "The id field needs to be a valid!")
	}

	if len(r.Name) > 50 {
		errs.Add("name", "The name field is longer then 50 characters long!")
	}

	if len(r.Description) > 100 {
		errs.Add("description", "The descripition field is longer then 100 characters long!")
	}

	return errs
}
