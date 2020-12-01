package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateKeypadConditionToLampEvent(keypadConditionsToLampEvents persistence.KeypadConditionsToLampEvents) int64 {
	rowsAffected := p.db.Create(&keypadConditionsToLampEvents).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetKeypadConditionToLampEvent(id string) (bool, persistence.KeypadConditionsToLampEvents) {
	var keypadConditionsToLampEvents persistence.KeypadConditionsToLampEvents
	recordNotFound := p.db.Where("id=?", id).First(&keypadConditionsToLampEvents).RecordNotFound()
	return recordNotFound, keypadConditionsToLampEvents
}

func (p MysqlPersistence) UpdateKeypadConditionToLampEvent(keypadConditionsToLampEvents persistence.KeypadConditionsToLampEvents) (bool, error) {
	recordNotFound := p.db.Where("id=?", keypadConditionsToLampEvents.ID).First(&persistence.KeypadConditionsToLampEvents{}).RecordNotFound()
	err := p.db.Model(&keypadConditionsToLampEvents).Where("id=?", keypadConditionsToLampEvents.ID).Updates(persistence.KeypadConditionsToLampEvents{ConditionID: keypadConditionsToLampEvents.ConditionID, EventID: keypadConditionsToLampEvents.EventID}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteKeypadConditionToLampEvent(keypadConditionsToLampEvents persistence.KeypadConditionsToLampEvents) (bool, error) {
	recordNotFound := p.db.Where("id=?", keypadConditionsToLampEvents.ID).First(&persistence.KeypadConditionsToLampEvents{}).RecordNotFound()
	err := p.db.Delete(&keypadConditionsToLampEvents).Error
	return recordNotFound, err
}
