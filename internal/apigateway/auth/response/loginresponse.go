package response

import "time"

// LoginRespose is the response that is returned when a user has successfully login.
// swagger:response LoginResponse
type LoginResponse struct {

	// The ID of the login user
	ID string `json:"id"`

	// The users username
	Username string `json:"username"`

	// The name of the user
	Name string `json:"name"`

	// The email of the user
	Email string `json:"email"`

	// The JWT token for the user
	Token string `json:"token"`

	// The expire time for the given JWT token
	Expires time.Time `json:"expired"`
}
