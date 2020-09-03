package request

import "net/url"

type GetUserRequest struct {
	Username string
}

func (r *GetUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	return errs
}
