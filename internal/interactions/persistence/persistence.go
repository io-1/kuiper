//go:generate mockgen -source persistence.go -destination=mock/mockpersistence.go -package=mock
package persistence

type Persistence interface {
	CreateInteraction(interaction Interaction) int64
	GetInteraction(id string) (bool, Interaction)
	UpdateInteraction(interaction Interaction) (bool, error)
	DeleteInteraction(interaction Interaction) (bool, error)
}
