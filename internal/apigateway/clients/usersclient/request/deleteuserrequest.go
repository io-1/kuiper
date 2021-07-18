package request

import "net/url"

// The request used to delete a user
type DeleteUserRequest struct{}

func (r *DeleteUserRequest) Validate(id string) url.Values {
	errs := url.Values{}
	// FIXME: uuids are in the form 8-4-4-4-12 for a total of 36 characters
	return errs
}
