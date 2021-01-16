package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampOnEvent(lampToggleEvent persistence.LampOnEvent) int64 {
	rowsAffected := p.db.Create(&lampToggleEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampOnEvent(id string) (recordNotFound bool, lampToggleEvent persistence.LampOnEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampToggleEvent).RecordNotFound()
	return recordNotFound, lampToggleEvent
}

func (p MysqlPersistence) UpdateLampOnEvent(lampToggleEvent persistence.LampOnEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampToggleEvent.ID).First(&persistence.LampOnEvent{}).RecordNotFound()
	err = p.db.Model(&lampToggleEvent).Where("id=?", lampToggleEvent.ID).Updates(persistence.LampOnEvent{Mac: lampToggleEvent.Mac}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampOnEvent(lampToggleEvent persistence.LampOnEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampToggleEvent.ID).First(&persistence.LampOnEvent{}).RecordNotFound()
	err = p.db.Delete(&lampToggleEvent).Error
	return recordNotFound, err
}
