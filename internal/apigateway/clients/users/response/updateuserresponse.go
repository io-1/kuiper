package response

// UpdateUserResponse is a response that ire returned when a user was successfully updated.
// swagger:response UpdateUserResponse
type UpdateUserResponse struct {

	// The ID of the user that is being updated.
	ID string `json:"id"`

	// The username of the user that is being updated.
	Username string `json:"username"`

	// The name of the user that is being updated.
	Name string `json:"name"`

	// The email of the user that is being updated.
	Email string `json:"email"`
}

// UpdateUserResponse is a response that ire returned when a user was successfully updated.
// swagger:response UpdateUserResponse
type UpdateUserResponseWrapper struct {
	// in: body
	Body UpdateUserResponse
}
