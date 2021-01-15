package request

import "net/url"

// CreateUserRequest is the request that is used for creating a user
type CreateUserRequest struct {

	// The username for the user
	Username string `json:"username" binding:"required"`

	// The password for the user
	Password string `json:"password" binding:"required"`

	// The name of the uesr
	Name string `json:"name" binding:"required"`

	// The email of the user
	Email string `json:"email" binding:"required"`
}

func (r *CreateUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	// FIXME: password 100 characters
	// FIXME: name 100 character
	// FIXME: email 100 characters
	// FIXME: check that email is valid
	return errs
}
