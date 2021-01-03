package persistence

type InteractionDetails struct {

	// FIXME: how to handle null values
	// Interaction Interaction

	// FIXME: how to handle null values
	KeypadCondition KeypadCondition

	// FIXME: how to handle null values
	// LampEvent LampEvent

	LampEventType string

	LampToggleEvent LampToggleEvent

	LampColorEvent LampColorEvent

	LampPulseEvent LampPulseEvent
}
