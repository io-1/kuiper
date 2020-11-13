package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateInteraction(interaction persistence.Interaction) int64 {
	return 0
}

func (p MysqlPersistence) GetInteraction(id string) (bool, persistence.Interaction) {
	return false, persistence.Interaction{}
}

func (p MysqlPersistence) UpdateInteraction(interaction persistence.Interaction) (bool, error) {
	return false, nil
}

func (p MysqlPersistence) DeleteInteraction(interaction persistence.Interaction) (bool, error) {
	return false, nil
}
