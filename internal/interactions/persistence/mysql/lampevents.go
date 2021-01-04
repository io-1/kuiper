package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateLampEvent(lampEvent persistence.LampEvent) int64 {
	rowsAffected := p.db.Create(&lampEvent).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetLampEvent(id string) (recordNotFound bool, lampEvent persistence.LampEvent) {
	recordNotFound = p.db.Where("id=?", id).First(&lampEvent).RecordNotFound()
	return recordNotFound, lampEvent
}

func (p MysqlPersistence) GetLampEventsByKeypadConditionID(id string) ([]persistence.LampEvent, error) {
	rows, err := p.db.Raw("select coalesce(lte.id, lce.id, lpe.id) as id, coalesce(lte.mac, lce.mac, lpe.mac) as mac, coalesce(lte.event_type, lce.event_type, lpe.event_type) as event_type, IFNULL(coalesce(lce.red, lpe.red),0) as red, IFNULL(coalesce(lce.green, lpe.green),0) as green, IFNULL(coalesce(lce.blue, lpe.blue),0) as blue, coalesce(lte.created_at, lce.created_at, lpe.created_at) as created_at, coalesce(lte.updated_at, lce.updated_at, lpe.updated_at) as updated_at, coalesce(lte.deleted_at, lce.deleted_at, lpe.deleted_at) as deleted_at from keypad_conditions_to_lamp_events ktl left join (select *, 'toggle' as event_type from lamp_toggle_events where deleted_at is null) lte on ktl.event_id = lte.id left join (select *, 'color' as event_type from lamp_color_events where deleted_at is null) lce on ktl.event_id = lce.id left join (select *, 'pulse' as event_type from lamp_pulse_events where deleted_at is null) lpe on ktl.event_id = lpe.id where ktl.condition_id = ?", id).Rows()
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
			&lampEvent.Red,
			&lampEvent.Green,
			&lampEvent.Blue,
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

func (p MysqlPersistence) GetLampEventsByKeypadMacAndButtonID(mac string, buttonID int) ([]persistence.LampEvent, error) {
	rows, err := p.db.Raw("select coalesce(lte.id, lce.id, lpe.id) as id, coalesce(lte.mac, lce.mac, lpe.mac) as mac, coalesce(lte.event_type, lce.event_type, lpe.event_type) as event_type, IFNULL(coalesce(lce.red, lpe.red),0) as red, IFNULL(coalesce(lce.green, lpe.green),0) as green, IFNULL(coalesce(lce.blue, lpe.blue),0) as blue, coalesce(lte.created_at, lce.created_at, lpe.created_at) as created_at, coalesce(lte.updated_at, lce.updated_at, lpe.updated_at) as updated_at, coalesce(lte.deleted_at, lce.deleted_at, lpe.deleted_at) as deleted_at from keypad_conditions k left join keypad_conditions_to_lamp_events ktl on k.id = ktl.condition_id left join (select *, 'toggle' as event_type from lamp_toggle_events where deleted_at is null) lte on ktl.event_id = lte.id left join (select *, 'color' as event_type from lamp_color_events where deleted_at is null) lce on ktl.event_id = lce.id left join (select *, 'pulse' as event_type from lamp_pulse_events where deleted_at is null) lpe on ktl.event_id = lpe.id where k.mac = ? and k.button_id = ?", mac, buttonID).Rows()
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
			&lampEvent.Red,
			&lampEvent.Green,
			&lampEvent.Blue,
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

func (p MysqlPersistence) UpdateLampEvent(lampEvent persistence.LampEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampEvent.ID).First(&persistence.LampEvent{}).RecordNotFound()
	err = p.db.Model(&lampEvent).Where("id=?", lampEvent.ID).Updates(persistence.LampEvent{Mac: lampEvent.Mac, EventType: lampEvent.EventType, Red: lampEvent.Red, Green: lampEvent.Green, Blue: lampEvent.Blue}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteLampEvent(lampEvent persistence.LampEvent) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", lampEvent.ID).First(&persistence.LampEvent{}).RecordNotFound()
	err = p.db.Delete(&lampEvent).Error
	return recordNotFound, err
}
