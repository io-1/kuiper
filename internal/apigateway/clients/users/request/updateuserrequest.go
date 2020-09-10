package request

import "net/url"

// The request used to update a user
type UpdateUserRequest struct {

	// The ID of the user being updated
	ID string `json:"id" binding:"required"`

	// The name of the user to update
	Name string `json:"name" binding:"required"`

	// The email of the user to update
	Email string `json:"email" binding:"required"`
}

// swagger:parameters UpdateUserRequest updateUser
type UpdateUserRequestWrapper struct {

	// The username of the user
	//
	// in: query
	Username string `json:"username"`

	// in: body
	Body UpdateUserRequest
}

func (r *UpdateUserRequest) Validate() url.Values {
	errs := url.Values{}
	// FIXME: password 100 characters
	// FIXME: name 100 character
	// FIXME: password 100 characters
	return errs
}
