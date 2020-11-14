package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateInteraction(interaction persistence.Interaction) int64 {
	rowsAffected := p.db.Create(&interaction).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetInteraction(id string) (bool, persistence.Interaction) {
	var interaction persistence.Interaction
	recordNotFound := p.db.Where("id=?", id).First(&interaction).RecordNotFound()
	return recordNotFound, interaction
}

func (p MysqlPersistence) UpdateInteraction(interaction persistence.Interaction) (bool, error) {
	recordNotFound := p.db.Where("id=?", interaction.ID).First(&persistence.Interaction{}).RecordNotFound()
	err := p.db.Model(&interaction).Where("id=?", interaction.ID).Updates(persistence.Interaction{Name: interaction.Name, Description: interaction.Description}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteInteraction(interaction persistence.Interaction) (bool, error) {
	recordNotFound := p.db.Where("id=?", interaction.ID).First(&persistence.Interaction{}).RecordNotFound()
	err := p.db.Delete(&interaction).Error
	return recordNotFound, err
}
