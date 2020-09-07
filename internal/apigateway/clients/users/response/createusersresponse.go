package response

// CreateUserResponse is a response when a user has been successfully created.
// swagger:response CreateUserResponse
type CreateUserResponse struct {

	// The ID of the user that was created.
	ID string `json:"id"`

	// The username of the user created.
	Username string `json:"username"`

	// The name of the user created.
	Name string `json:"name"`

	// The email of the user created.
	Email string `json:"email"`
}
