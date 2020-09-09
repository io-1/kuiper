package request

import "net/url"

// The request used to update a user
// swagger:parameters UpdateUserRequest updateUser
type UpdateUserRequest struct {

	// The ID of the user being updated
	ID       string `json:"id" binding:"required"`
	Username string `json:"-"`

	// The name of the user to update
	Name string `json:"name" binding:"required"`

	// The email of the user to update
	Email string `json:"email" binding:"required"`
}

func (r *UpdateUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: password 100 characters
	// FIXME: name 100 character
	// FIXME: password 100 characters
	return errs
}
