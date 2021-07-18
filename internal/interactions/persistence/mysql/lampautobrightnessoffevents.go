package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampAutoBrightnessOffEvent(lampAutoBrightnessOffEvent persistence.LampAutoBrightnessOffEvent) int64 {
	rowsAffected := p.db.Create(&lampAutoBrightnessOffEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampAutoBrightnessOffEvent(id string) (recordNotFound bool, lampAutoBrightnessOffEvent persistence.LampAutoBrightnessOffEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampAutoBrightnessOffEvent).RecordNotFound()
	return recordNotFound, lampAutoBrightnessOffEvent
}

func (p MysqlPersistence) UpdateLampAutoBrightnessOffEvent(lampAutoBrightnessOffEvent persistence.LampAutoBrightnessOffEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampAutoBrightnessOffEvent.ID).First(&persistence.LampAutoBrightnessOffEvent{}).RecordNotFound()
	err = p.db.Model(&lampAutoBrightnessOffEvent).Where("id=?", lampAutoBrightnessOffEvent.ID).Updates(persistence.LampAutoBrightnessOffEvent{Mac: lampAutoBrightnessOffEvent.Mac}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampAutoBrightnessOffEvent(lampAutoBrightnessOffEvent persistence.LampAutoBrightnessOffEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampAutoBrightnessOffEvent.ID).First(&persistence.LampAutoBrightnessOffEvent{}).RecordNotFound()
	err = p.db.Delete(&lampAutoBrightnessOffEvent).Error
	return recordNotFound, err
}
