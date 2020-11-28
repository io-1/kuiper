package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampEvent(keypadCondition persistence.LampEvent) int64 {
	rowsAffected := p.db.Create(&keypadCondition).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampEvent(id string) (bool, persistence.LampEvent) {
	var lampEvent persistence.LampEvent
	recordNotFound := p.db.Where("id=?", id).First(&lampEvent).RecordNotFound()
	return recordNotFound, lampEvent
}

func (p MysqlPersistence) UpdateLampEvent(lampEvent persistence.LampEvent) (bool, error) {
	recordNotFound := p.db.Where("id=?", lampEvent.ID).First(&persistence.LampEvent{}).RecordNotFound()
	err := p.db.Model(&lampEvent).Where("id=?", lampEvent.ID).Updates(persistence.LampEvent{Mac: lampEvent.Mac, EventType: lampEvent.EventType, Color: lampEvent.Color}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampEvent(lampEvent persistence.LampEvent) (bool, error) {
	recordNotFound := p.db.Where("id=?", lampEvent.ID).First(&persistence.LampEvent{}).RecordNotFound()
	err := p.db.Delete(&lampEvent).Error
	return recordNotFound, err
}
