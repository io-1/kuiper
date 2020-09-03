package request

import "net/url"

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (r *CreateUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: username 50 characters
	// FIXME: password 100 characters
	// FIXME: name 100 character
	// FIXME: password 100 characters
	return errs
}
