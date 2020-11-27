package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateKeypadCondition(keypadCondition persistence.KeypadCondition) int64 {
	rowsAffected := p.db.Create(&keypadCondition).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetKeypadCondition(id string) (bool, persistence.KeypadCondition) {
	var keypadCondition persistence.KeypadCondition
	recordNotFound := p.db.Where("id=?", id).First(&keypadCondition).RecordNotFound()
	return recordNotFound, keypadCondition
}

func (p MysqlPersistence) UpdateKeypadCondition(keypadCondition persistence.KeypadCondition) (bool, error) {
	recordNotFound := p.db.Where("id=?", keypadCondition.ID).First(&persistence.KeypadCondition{}).RecordNotFound()
	err := p.db.Model(&keypadCondition).Where("id=?", keypadCondition.ID).Updates(persistence.KeypadCondition{InteractionID: keypadCondition.InteractionID, Mac: keypadCondition.Mac, ButtonID: keypadCondition.ButtonID}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteKeypadCondition(keypadCondition persistence.KeypadCondition) (bool, error) {
	recordNotFound := p.db.Where("id=?", keypadCondition.ID).First(&persistence.KeypadCondition{}).RecordNotFound()
	err := p.db.Delete(&keypadCondition).Error
	return recordNotFound, err
}
