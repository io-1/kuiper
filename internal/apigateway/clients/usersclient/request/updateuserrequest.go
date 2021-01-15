package request

import "net/url"

// The request used to update a user
type UpdateUserRequest struct {

	// The username of the user to update
	Username string `json:"username" binding:"required"`

	// The name of the user to update
	Name string `json:"name" binding:"required"`

	// The email of the user to update
	Email string `json:"email" binding:"required"`
}

func (r *UpdateUserRequest) Validate(id string) url.Values {
	errs := url.Values{}
	// FIXME: uuids are in the form 8-4-4-4-12 for a total of 36 characters
	// FIXME: username 50 character
	// FIXME: name 100 character
	// FIXME: check email
	return errs
}
