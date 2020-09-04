package response

import "time"

type LoginResponse struct {
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Token    string    `json:"token"`
	Expires  time.Time `json:"expired"`
}
