package response

// Updates a user.
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
