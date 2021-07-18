package request

import (
	"net/url"
	"regexp"
)

// PatchUserRequest is the request that is used for patching a user
type PatchUserRequest struct {

	// The username for the user
	Username string `json:"username"`

	// The name of the uesr
	Name string `json:"name"`

	// The email of the user
	Email string `json:"email"`
}

func (r *PatchUserRequest) Validate(id string) url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	validID := regex.MatchString(id)
	if !validID {
		errs.Add("id", "The id field needs to be a valid!")
	}

	// username 50 characters
	if len(r.Username) > 50 {
		errs.Add("username", "The username field needs to be less then 50 characters")
	}

	// name 100 character
	if len(r.Name) > 100 {
		errs.Add("name", "The name field needs to be less then 100 characters")
	}

	// email 100 characters
	if len(r.Email) > 100 {
		errs.Add("email", "The email field needs to be less then 100 characters")
	}

	// FIXME: check that email is valid
	// if r.Email != "" {
	// }

	return errs
}
