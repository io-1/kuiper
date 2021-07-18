package response

type GetInteractionDetailsResponse struct {
	ID           string                                    `json:"id"`
	Interactions []KeypadConditionsToLampEventsInteraction `json:"interactions"`
}

type KeypadConditionsToLampEventsInteraction struct {
	KeypadCondition KeypadCondition `json:"keypadCondition"`
	LampEvent       interface{}     `json:"lampEvent"`
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

type LampOnEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
}

type LampOffEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
}

type LampToggleEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
}

type LampBrightnessEvent struct {
	ID         string `json:"id"`
	Mac        string `json:"mac"`
	Brightness int32  `json:"brightness"`
	EventType  string `json:"eventType"`
}

type LampAutoBrightnessOnEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
}

type LampAutoBrightnessOffEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
}

type LampAutoBrightnessToggleEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
}

type LampColorEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
	Red       int32  `json:"red"`
	Green     int32  `json:"green"`
	Blue      int32  `json:"blue"`
}

type LampPulseEvent struct {
	ID        string `json:"id"`
	Mac       string `json:"mac"`
	EventType string `json:"eventType"`
	Red       int32  `json:"red"`
	Green     int32  `json:"green"`
	Blue      int32  `json:"blue"`
}
