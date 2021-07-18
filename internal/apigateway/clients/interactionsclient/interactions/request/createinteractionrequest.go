package request

import "net/url"

type CreateInteractionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (r CreateInteractionRequest) Validate() url.Values {
	errs := url.Values{}

	// name length 50 characters
	if len(r.Name) > 50 {
		errs.Add("name", "The name field is longer then 50 characters long!")
	}

	// description length 100 characters
	if len(r.Description) > 100 {
		errs.Add("descripttion", "The description field is longer then 100 characters long!")
	}

	return errs
}
