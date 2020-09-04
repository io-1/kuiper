package request

import "net/url"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *LoginRequest) Validate() url.Values {
	errs := url.Values{}

	if len(r.Username) > 50 {
		errs.Add("username", "longer then 50 characters!")
	}

	if len(r.Password) > 100 {
		errs.Add("password", "longer then 100 characters!")
	}

	return errs
}
