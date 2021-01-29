package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampBrightnessEvent(lampBrightnessEvent persistence.LampBrightnessEvent) int64 {
	rowsAffected := p.db.Create(&lampBrightnessEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampBrightnessEvent(id string) (recordNotFound bool, lampBrightnessEvent persistence.LampBrightnessEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampBrightnessEvent).RecordNotFound()
	return recordNotFound, lampBrightnessEvent
}

func (p MysqlPersistence) UpdateLampBrightnessEvent(lampBrightnessEvent persistence.LampBrightnessEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampBrightnessEvent.ID).First(&persistence.LampBrightnessEvent{}).RecordNotFound()
	err = p.db.Model(&lampBrightnessEvent).Where("id=?", lampBrightnessEvent.ID).Updates(persistence.LampBrightnessEvent{Mac: lampBrightnessEvent.Mac, Brightness: lampBrightnessEvent.Brightness}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampBrightnessEvent(lampBrightnessEvent persistence.LampBrightnessEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampBrightnessEvent.ID).First(&persistence.LampBrightnessEvent{}).RecordNotFound()
	err = p.db.Delete(&lampBrightnessEvent).Error
	return recordNotFound, err
}
