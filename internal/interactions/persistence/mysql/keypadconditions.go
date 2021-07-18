package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateKeypadCondition(keypadCondition persistence.KeypadCondition) int64 {
	rowsAffected := p.db.Create(&keypadCondition).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetKeypadCondition(id string) (recordNotFound bool, keypadCondition persistence.KeypadCondition) {
	recordNotFound = p.db.Where("id=?", id).First(&keypadCondition).RecordNotFound()
	return recordNotFound, keypadCondition
}

func (p MysqlPersistence) GetKeypadConditionByMac(mac string) (recordNotFound bool, keypadCondition persistence.KeypadCondition) {
	recordNotFound = p.db.Where("mac=?", mac).First(&keypadCondition).RecordNotFound()
	return recordNotFound, keypadCondition
}

func (p MysqlPersistence) GetKeypadConditionByMacAndButtonID(mac string, buttonID int) (recordNotFound bool, keypadCondition persistence.KeypadCondition) {
	recordNotFound = p.db.Where("mac=? AND button_id=?", mac, buttonID).First(&keypadCondition).RecordNotFound()
	return recordNotFound, keypadCondition
}

func (p MysqlPersistence) UpdateKeypadCondition(keypadCondition persistence.KeypadCondition) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", keypadCondition.ID).First(&persistence.KeypadCondition{}).RecordNotFound()
	err = p.db.Model(&keypadCondition).Where("id=?", keypadCondition.ID).Updates(persistence.KeypadCondition{Mac: keypadCondition.Mac, ButtonID: keypadCondition.ButtonID}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteKeypadCondition(keypadCondition persistence.KeypadCondition) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", keypadCondition.ID).First(&persistence.KeypadCondition{}).RecordNotFound()
	err = p.db.Delete(&keypadCondition).Error
	return recordNotFound, err
}
