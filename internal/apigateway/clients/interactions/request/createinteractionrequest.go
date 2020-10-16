package request

import "net/url"

type CreateInteractionConditionRequest struct {
	Mac                 string `json:"mac"`
	Sensor              string `json:"sensor"`
	Measurement         string `json:"measurement"`
	MeasurementOperator string `json:"measurement_operator"`
	MeasurementValue    string `json:"measurement_value"`
}

type CreateInteractionActionRequest struct {
	Mac    string `json:"mac"`
	Action string `json:"action"`
}

type CreateInteractionRequest struct {
	Name       string                            `json:"name"`
	Conditions CreateInteractionConditionRequest `json:"conditions"`
	Actions    CreateInteractionActionRequest    `json:"actions"`
}

func (r CreateInteractionRequest) Validate() url.Values {
	errs := url.Values{}
	return errs
}
