package request

import "net/url"

// Request used to get a user
// swagger:parameters getUser
type GetUserRequest struct {
	Username string `json:"-"`
}

func (r *GetUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	return errs
}
