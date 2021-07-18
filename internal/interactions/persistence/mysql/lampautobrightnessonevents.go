package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampAutoBrightnessOnEvent(lampAutoBrightnessOnEvent persistence.LampAutoBrightnessOnEvent) int64 {
	rowsAffected := p.db.Create(&lampAutoBrightnessOnEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampAutoBrightnessOnEvent(id string) (recordNotFound bool, lampAutoBrightnessOnEvent persistence.LampAutoBrightnessOnEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampAutoBrightnessOnEvent).RecordNotFound()
	return recordNotFound, lampAutoBrightnessOnEvent
}

func (p MysqlPersistence) UpdateLampAutoBrightnessOnEvent(lampAutoBrightnessOnEvent persistence.LampAutoBrightnessOnEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampAutoBrightnessOnEvent.ID).First(&persistence.LampAutoBrightnessOnEvent{}).RecordNotFound()
	err = p.db.Model(&lampAutoBrightnessOnEvent).Where("id=?", lampAutoBrightnessOnEvent.ID).Updates(persistence.LampAutoBrightnessOnEvent{Mac: lampAutoBrightnessOnEvent.Mac}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampAutoBrightnessOnEvent(lampAutoBrightnessOnEvent persistence.LampAutoBrightnessOnEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampAutoBrightnessOnEvent.ID).First(&persistence.LampAutoBrightnessOnEvent{}).RecordNotFound()
	err = p.db.Delete(&lampAutoBrightnessOnEvent).Error
	return recordNotFound, err
}
