package request

import "net/url"

// Request used to get a user
type GetUserRequest struct{}

func (r *GetUserRequest) Validate(username string) url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	return errs
}
