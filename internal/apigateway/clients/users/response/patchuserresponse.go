package response

// PatchUserResponse is a response when a user has been successfully patched.
type PatchUserResponse struct {

	// The ID of the user that was patched.
	ID string `json:"id"`

	// The username of the user patched.
	Username string `json:"username"`

	// The name of the user patched.
	Name string `json:"name"`

	// The email of the user patched.
	Email string `json:"email"`
}
