package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateConditionsToEvents(conditionsToEvents persistence.ConditionsToEvents) int64 {
	rowsAffected := p.db.Create(&conditionsToEvents).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetConditionsToEvents(id string) (bool, persistence.ConditionsToEvents) {
	var conditionsToEvents persistence.ConditionsToEvents
	recordNotFound := p.db.Where("id=?", id).First(&conditionsToEvents).RecordNotFound()
	return recordNotFound, conditionsToEvents
}

func (p MysqlPersistence) UpdateConditionsToEvents(conditionsToEvents persistence.ConditionsToEvents) (bool, error) {
	recordNotFound := p.db.Where("id=?", conditionsToEvents.ID).First(&persistence.ConditionsToEvents{}).RecordNotFound()
	err := p.db.Model(&conditionsToEvents).Where("id=?", conditionsToEvents.ID).Updates(persistence.ConditionsToEvents{InteractionID: conditionsToEvents.InteractionID, ConditionID: conditionsToEvents.ConditionID, EventID: conditionsToEvents.EventID}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteConditionsToEvents(conditionsToEvents persistence.ConditionsToEvents) (bool, error) {
	recordNotFound := p.db.Where("id=?", conditionsToEvents.ID).First(&persistence.ConditionsToEvents{}).RecordNotFound()
	err := p.db.Delete(&conditionsToEvents).Error
	return recordNotFound, err
}
