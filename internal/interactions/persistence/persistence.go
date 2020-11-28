//go:generate mockgen -source persistence.go -destination=mock/mockpersistence.go -package=mock
package persistence

type Persistence interface {
	CreateInteraction(interaction Interaction) int64
	GetInteraction(id string) (bool, Interaction)
	UpdateInteraction(interaction Interaction) (bool, error)
	DeleteInteraction(interaction Interaction) (bool, error)

	CreateKeypadCondition(keypadCondition KeypadCondition) int64
	GetKeypadCondition(id string) (bool, KeypadCondition)
	UpdateKeypadCondition(keypadCondition KeypadCondition) (bool, error)
	DeleteKeypadCondition(keypadCondition KeypadCondition) (bool, error)

	CreateLampEvent(keypadCondition LampEvent) int64
	GetLampEvent(id string) (bool, LampEvent)
	UpdateLampEvent(lampEvent LampEvent) (bool, error)
	DeleteLampEvent(lampEvent LampEvent) (bool, error)

	CreateConditionsToEvents(conditionsToEvents ConditionsToEvents) int64
	GetConditionsToEvents(id string) (bool, ConditionsToEvents)
	UpdateConditionsToEvents(conditionsToEvents ConditionsToEvents) (bool, error)
	DeleteConditionsToEvents(conditionsToEvents ConditionsToEvents) (bool, error)
}
