package request

import "net/url"

// The request used to delete a user
// swagger:parameters deleteUser
type DeleteUserRequest struct {

	// The username of the user
	//
	// in: query
	Username string `json:"username"`
}

func (r *DeleteUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	return errs
}
