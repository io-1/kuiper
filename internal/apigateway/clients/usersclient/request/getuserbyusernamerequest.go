package request

import "net/url"

type GetUserByUsernameRequest struct{}

func (r GetUserByUsernameRequest) Validate(username string) url.Values {
	errs := url.Values{}

	if len(username) > 50 {
		errs.Add("username", "longer then 50 characters!")
	}

	return errs
}
