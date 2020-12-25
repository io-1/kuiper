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
	UpdateKeypadCondition(keypadCondition KeypadCondition) (recordNotFound bool, err error)
	DeleteKeypadCondition(keypadCondition KeypadCondition) (recordNotFound bool, err error)

	CreateLampEvent(keypadCondition LampEvent) int64
	GetLampEvent(id string) (recordNotFound bool, lampEvent LampEvent)
	GetLampEventsByKeypadConditionID(id string) ([]LampEvent, error)
	UpdateLampEvent(lampEvent LampEvent) (recordNotFound bool, err error)
	DeleteLampEvent(lampEvent LampEvent) (recordNotFound bool, err error)

	CreateConditionsToEvents(conditionsToEvents ConditionsToEvents) int64
	GetConditionsToEvents(id string) (recordNotFound bool, conditionsToEvents ConditionsToEvents)
	UpdateConditionsToEvents(conditionsToEvents ConditionsToEvents) (recordNotFound bool, err error)
	DeleteConditionsToEvents(conditionsToEvents ConditionsToEvents) (recordNotFound bool, err error)
}
