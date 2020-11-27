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
}
