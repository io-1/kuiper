package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampAutoBrightnessToggleEvent(lampAutoBrightnessToggleEvent persistence.LampAutoBrightnessToggleEvent) int64 {
	rowsAffected := p.db.Create(&lampAutoBrightnessToggleEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampAutoBrightnessToggleEvent(id string) (recordNotFound bool, lampAutoBrightnessToggleEvent persistence.LampAutoBrightnessToggleEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampAutoBrightnessToggleEvent).RecordNotFound()
	return recordNotFound, lampAutoBrightnessToggleEvent
}

func (p MysqlPersistence) UpdateLampAutoBrightnessToggleEvent(lampAutoBrightnessToggleEvent persistence.LampAutoBrightnessToggleEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampAutoBrightnessToggleEvent.ID).First(&persistence.LampAutoBrightnessToggleEvent{}).RecordNotFound()
	err = p.db.Model(&lampAutoBrightnessToggleEvent).Where("id=?", lampAutoBrightnessToggleEvent.ID).Updates(persistence.LampAutoBrightnessToggleEvent{Mac: lampAutoBrightnessToggleEvent.Mac}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampAutoBrightnessToggleEvent(lampAutoBrightnessToggleEvent persistence.LampAutoBrightnessToggleEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampAutoBrightnessToggleEvent.ID).First(&persistence.LampAutoBrightnessToggleEvent{}).RecordNotFound()
	err = p.db.Delete(&lampAutoBrightnessToggleEvent).Error
	return recordNotFound, err
}
