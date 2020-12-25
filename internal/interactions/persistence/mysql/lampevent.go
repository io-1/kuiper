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

func (p MysqlPersistence) GetLampEventsByKeypadConditionID(id string) ([]persistence.LampEvent, error) {
	rows, err := p.db.Table("lamp_events").Select("lamp_events.id, lamp_events.mac, lamp_events.event_type, lamp_events.color, lamp_events.created_at, lamp_events.updated_at, lamp_events.deleted_at").Joins("left join conditions_to_events on conditions_to_events.event_id = lamp_events.id").Where("conditions_to_events.condition_id=?", id).Rows()
	if err != nil {
		return []persistence.LampEvent{}, err
	}

	var allLampEvents []persistence.LampEvent
	for rows.Next() {
		var lampEvent persistence.LampEvent
		err = rows.Scan(
			&lampEvent.ID,
			&lampEvent.Mac,
			&lampEvent.EventType,
			&lampEvent.Color,
			&lampEvent.CreatedAt,
			&lampEvent.UpdatedAt,
			&lampEvent.DeletedAt,
		)

		if err != nil {
			return []persistence.LampEvent{}, err
		}

		allLampEvents = append(allLampEvents, lampEvent)
	}

	return allLampEvents, nil
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
