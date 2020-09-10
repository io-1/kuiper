package response

// DeleteUserResponse is a response when a user has been successfully deleted.
// swagger:response DeleteUserResponse
type DeleteUserResponse struct {

	// The username of the user that was deleted.
	Username string `json:"username"`
}

// DeleteUserResponse is a response when a user has been successfully deleted.
// swagger:response DeleteUserResponse
type DeleteUserResponseWrapper struct {
	// in: body
	Body DeleteUserResponse
}
