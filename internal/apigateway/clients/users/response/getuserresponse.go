package response

// GetUserResponse is a response when getting a user successfully.
// swagger:response GetUserResponse
type GetUserResponse struct {

	// The ID of the user.
	ID string `json:"id"`

	// The username of the user.
	Username string `json:"username"`

	// The name of the user.
	Name string `json:"name"`

	// The email of the user.
	Email string `json:"email"`
}
