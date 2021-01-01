package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampToggleEvent(lampToggleEvent persistence.LampToggleEvent) int64 {
	rowsAffected := p.db.Create(&lampToggleEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampToggleEvent(id string) (recordNotFound bool, lampToggleEvent persistence.LampToggleEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampToggleEvent).RecordNotFound()
	return recordNotFound, lampToggleEvent
}

func (p MysqlPersistence) UpdateLampToggleEvent(lampToggleEvent persistence.LampToggleEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampToggleEvent.ID).First(&persistence.LampToggleEvent{}).RecordNotFound()
	err = p.db.Model(&lampToggleEvent).Where("id=?", lampToggleEvent.ID).Updates(persistence.LampToggleEvent{Mac: lampToggleEvent.Mac}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampToggleEvent(lampToggleEvent persistence.LampToggleEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampToggleEvent.ID).First(&persistence.LampToggleEvent{}).RecordNotFound()
	err = p.db.Delete(&lampToggleEvent).Error
	return recordNotFound, err
}
