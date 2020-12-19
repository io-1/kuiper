package response

type GetInteractionDetailsResponse struct {
	ID string `json:"id"`
	// Name         string                                    `json:"name"`
	// Description  string                                    `json:"description"`
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
	Color     string `json:"color"`
}
