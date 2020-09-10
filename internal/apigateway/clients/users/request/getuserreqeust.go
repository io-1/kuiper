package request

import "net/url"

// Request used to get a user
// swagger:parameters getUser
type GetUserRequest struct {

	// The username of the user
	//
	// in: query
	Username string `json:"username"`
}

func (r *GetUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	return errs
}
