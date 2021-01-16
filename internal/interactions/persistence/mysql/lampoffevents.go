package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampOffEvent(lampToggleEvent persistence.LampOffEvent) int64 {
	rowsAffected := p.db.Create(&lampToggleEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampOffEvent(id string) (recordNotFound bool, lampToggleEvent persistence.LampOffEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampToggleEvent).RecordNotFound()
	return recordNotFound, lampToggleEvent
}

func (p MysqlPersistence) UpdateLampOffEvent(lampToggleEvent persistence.LampOffEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampToggleEvent.ID).First(&persistence.LampOffEvent{}).RecordNotFound()
	err = p.db.Model(&lampToggleEvent).Where("id=?", lampToggleEvent.ID).Updates(persistence.LampOffEvent{Mac: lampToggleEvent.Mac}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampOffEvent(lampToggleEvent persistence.LampOffEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampToggleEvent.ID).First(&persistence.LampOffEvent{}).RecordNotFound()
	err = p.db.Delete(&lampToggleEvent).Error
	return recordNotFound, err
}
