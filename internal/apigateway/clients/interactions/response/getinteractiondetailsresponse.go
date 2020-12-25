package response

type GetInteractionDetailsResponse struct {
	ID           string                                    `json:"id"`
	Interactions []KeypadConditionsToLampEventsInteraction `json:"interactions"`
}

type KeypadConditionsToLampEventsInteraction struct {
	KeypadCondition KeypadCondition `json:"keypadCondition"`
	LampEvent       LampEvent       `json:"lampEvent"`
}

type KeypadCondition struct {
	ID       string `json:"id"`
	Mac      string `json:"mac"`
	ButtonID int32  `json:"buttonID"`
}

type LampEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
	Red       int32  `json:"red"`
	Green     int32  `json:"green"`
	Blue      int32  `json:"blue"`
}
