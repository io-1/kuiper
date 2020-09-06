package request

import "net/url"

type UpdateUserRequest struct {
	ID       string `json:"id" binding:"required"`
	Username string `json:"-"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (r *UpdateUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: password 100 characters
	// FIXME: name 100 character
	// FIXME: password 100 characters
	return errs
}
