package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampPulseEvent(lampPulseEvent persistence.LampPulseEvent) int64 {
	rowsAffected := p.db.Create(&lampPulseEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampPulseEvent(id string) (recordNotFound bool, lampPulseEvent persistence.LampPulseEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampPulseEvent).RecordNotFound()
	return recordNotFound, lampPulseEvent
}

func (p MysqlPersistence) UpdateLampPulseEvent(lampPulseEvent persistence.LampPulseEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampPulseEvent.ID).First(&persistence.LampPulseEvent{}).RecordNotFound()
	err = p.db.Model(&lampPulseEvent).Where("id=?", lampPulseEvent.ID).Updates(persistence.LampPulseEvent{Mac: lampPulseEvent.Mac, Red: lampPulseEvent.Red, Green: lampPulseEvent.Green, Blue: lampPulseEvent.Blue}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampPulseEvent(lampPulseEvent persistence.LampPulseEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampPulseEvent.ID).First(&persistence.LampPulseEvent{}).RecordNotFound()
	err = p.db.Delete(&lampPulseEvent).Error
	return recordNotFound, err
}
