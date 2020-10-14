package response

type CreateInteractionConditionResponse struct {
	ID                  string `json:"id"`
	Mac                 string `json:"mac"`
	Sensor              string `json:"sensor"`
	Measurement         string `json:"measurement"`
	MeasurementOperator string `json:"measurement_operator"`
	MeasurementValue    string `json:"measurement_value"`
}

type CreateInteractionActionResponse struct {
	ID     string `json:"id"`
	Mac    string `json:"mac"`
	Action string `json:"action"`
}

type CreateInteractionResponse struct {
	ID         string                               `json:"id"`
	Name       string                               `json:"name"`
	Conditions []CreateInteractionConditionResponse `json:"conditions"`
	Actions    []CreateInteractionActionResponse    `json:"actions"`
}
