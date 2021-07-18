package request

import "net/url"

// Request used to get a user
type GetUserRequest struct{}

func (r *GetUserRequest) Validate(id string) url.Values {
	errs := url.Values{}
	// FIXME: uuids are in the form 8-4-4-4-12 for a total of 36 characters
	return errs
}
