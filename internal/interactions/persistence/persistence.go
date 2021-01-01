//go:generate mockgen -source persistence.go -destination=mock/mockpersistence.go -package=mock
package persistence

type Persistence interface {
	CreateInteraction(interaction Interaction) int64
	GetInteraction(id string) (recordNotFound bool, interaction Interaction)
	GetInteractionDetails(id string) ([]InteractionDetails, error)
	UpdateInteraction(interaction Interaction) (recordNotFound bool, err error)
	DeleteInteraction(interaction Interaction) (recordNotFound bool, err error)

	CreateKeypadCondition(keypadCondition KeypadCondition) int64
	GetKeypadCondition(id string) (recordNotFound bool, keypadCondition KeypadCondition)
	GetKeypadConditionByMac(mac string) (recordNotFound bool, keypadCondition KeypadCondition)
	GetKeypadConditionByMacAndButtonID(mac string, buttonID int) (recordNotFound bool, keypadCondition KeypadCondition)
	UpdateKeypadCondition(keypadCondition KeypadCondition) (recordNotFound bool, err error)
	DeleteKeypadCondition(keypadCondition KeypadCondition) (recordNotFound bool, err error)

	CreateLampEvent(lampEvent LampEvent) int64
	GetLampEvent(id string) (recordNotFound bool, lampEvent LampEvent)
	GetLampEventsByKeypadConditionID(id string) ([]LampEvent, error)
	UpdateLampEvent(lampEvent LampEvent) (recordNotFound bool, err error)
	DeleteLampEvent(lampEvent LampEvent) (recordNotFound bool, err error)

	CreateLampToggleEvent(lampToggleEvent LampToggleEvent) int64
	GetLampToggleEvent(id string) (recordNotFound bool, lampToggleEvent LampToggleEvent)
	UpdateLampToggleEvent(lampToggleEvent LampToggleEvent) (recordNotFound bool, err error)
	DeleteLampToggleEvent(lampToggleEvent LampToggleEvent) (recordNotFound bool, err error)

	CreateLampColorEvent(lampColorEvent LampColorEvent) int64
	GetLampColorEvent(id string) (recordNotFound bool, lampColorEvent LampColorEvent)
	UpdateLampColorEvent(lampColorEvent LampColorEvent) (recordNotFound bool, err error)
	DeleteLampColorEvent(lampColorEvent LampColorEvent) (recordNotFound bool, err error)

	CreateLampPulseEvent(lampPulseEvent LampPulseEvent) int64
	GetLampPulseEvent(id string) (recordNotFound bool, lampPulseEvent LampPulseEvent)
	UpdateLampPulseEvent(lampPulseEvent LampPulseEvent) (recordNotFound bool, err error)
	DeleteLampPulseEvent(lampPulseEvent LampPulseEvent) (recordNotFound bool, err error)

	CreateConditionsToEvents(conditionsToEvents ConditionsToEvents) int64
	GetConditionsToEvents(id string) (recordNotFound bool, conditionsToEvents ConditionsToEvents)
	UpdateConditionsToEvents(conditionsToEvents ConditionsToEvents) (recordNotFound bool, err error)
	DeleteConditionsToEvents(conditionsToEvents ConditionsToEvents) (recordNotFound bool, err error)

	CreateKeypadConditionsToLampEvents(conditionsToEvents KeypadConditionsToLampEvents) int64
	GetKeypadConditionsToLampEvents(id string) (recordNotFound bool, conditionsToEvents KeypadConditionsToLampEvents)
	UpdateKeypadConditionsToLampEvents(conditionsToEvents KeypadConditionsToLampEvents) (recordNotFound bool, err error)
	DeleteKeypadConditionsToLampEvents(conditionsToEvents KeypadConditionsToLampEvents) (recordNotFound bool, err error)
}
