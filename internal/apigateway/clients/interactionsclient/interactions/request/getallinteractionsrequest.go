package request

import (
	"net/url"
)

type GetAllInteractionsRequest struct {
	Limit  *int32 `form:"limit" binding:"required"`
	Offset *int32 `form:"offset" binding:"required"`
}

func (r GetAllInteractionsRequest) Validate() url.Values {
	errs := url.Values{}

	if *r.Limit < 0 {
		errs.Add("limit", "The limit field needs to be greater then or equal to 0!")
	}

	if *r.Limit > 100 {
		errs.Add("limit", "The limit field needs to be greater then or equal to 100!")
	}

	if *r.Offset < 0 {
		errs.Add("offset", "The offset field needs to be greater then or equal to 0!")
	}

	if *r.Offset > 100 {
		errs.Add("offset", "The offset field needs to be greater then or equal to 100!")
	}

	return errs
}
