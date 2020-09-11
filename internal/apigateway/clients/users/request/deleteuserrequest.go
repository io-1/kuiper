package request

import "net/url"

// The request used to delete a user
type DeleteUserRequest struct{}

func (r *DeleteUserRequest) Validate(username string) url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	return errs
}
