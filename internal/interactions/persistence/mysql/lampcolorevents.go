package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampColorEvent(lampColorEvent persistence.LampColorEvent) int64 {
	rowsAffected := p.db.Create(&lampColorEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampColorEvent(id string) (recordNotFound bool, lampColorEvent persistence.LampColorEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampColorEvent).RecordNotFound()
	return recordNotFound, lampColorEvent
}

func (p MysqlPersistence) UpdateLampColorEvent(lampColorEvent persistence.LampColorEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampColorEvent.ID).First(&persistence.LampColorEvent{}).RecordNotFound()
	err = p.db.Model(&lampColorEvent).Where("id=?", lampColorEvent.ID).Updates(persistence.LampColorEvent{Mac: lampColorEvent.Mac, Red: lampColorEvent.Red, Green: lampColorEvent.Green, Blue: lampColorEvent.Blue}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampColorEvent(lampColorEvent persistence.LampColorEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampColorEvent.ID).First(&persistence.LampColorEvent{}).RecordNotFound()
	err = p.db.Delete(&lampColorEvent).Error
	return recordNotFound, err
}
