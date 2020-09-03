package request

import "net/url"

type DeleteUserRequest struct {
	Username string
}

func (r *DeleteUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	return errs
}
