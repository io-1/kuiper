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
	query := `SELECT 
		COALESCE(lte.id, lce.id, lpe.id) AS id, 
		COALESCE(lte.mac, lce.mac, lpe.mac) AS mac, 
		COALESCE(lte.event_type, lce.event_type, lpe.event_type) AS event_type, 
		IFNULL(COALESCE(lce.red, lpe.red),0) AS red, 
		IFNULL(COALESCE(lce.green, lpe.green),0) AS green, 
		IFNULL(COALESCE(lce.blue, lpe.blue),0) AS blue, 
		COALESCE(lte.created_at, lce.created_at, lpe.created_at) AS created_at, 
		COALESCE(lte.updated_at, lce.updated_at, lpe.updated_at) AS updated_at, 
		COALESCE(lte.deleted_at, lce.deleted_at, lpe.deleted_at) AS deleted_at 
	FROM keypad_conditions_to_lamp_events ktl 
		left join 
			(SELECT 
				*, 
				'toggle' AS event_type 
			FROM lamp_toggle_events WHERE deleted_at IS NULL) lte ON ktl.event_id = lte.id 
		left join 
			(SELECT 
				*, 
				'color' AS event_type 
			FROM lamp_color_events WHERE deleted_at IS NULL) lce ON ktl.event_id = lce.id 
		left join 
			(SELECT 
				*, 
				'pulse' AS event_type 
			FROM lamp_pulse_events WHERE deleted_at IS NULL) lpe ON ktl.event_id = lpe.id 
	WHERE ktl.condition_id = ?`

	rows, err := p.db.Raw(query, id).Rows()
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
	query := `SELECT 
		COALESCE(lte.id, lce.id, lpe.id) AS id, 
		COALESCE(lte.mac, lce.mac, lpe.mac) AS mac, 
		COALESCE(lte.event_type, lce.event_type, lpe.event_type) AS event_type, 
		IFNULL(COALESCE(lce.red, lpe.red),0) AS red, 
		IFNULL(COALESCE(lce.green, lpe.green),0) AS green, 
		IFNULL(COALESCE(lce.blue, lpe.blue),0) AS blue, 
		COALESCE(lte.created_at, lce.created_at, lpe.created_at) AS created_at, 
		COALESCE(lte.updated_at, lce.updated_at, lpe.updated_at) AS updated_at, 
		COALESCE(lte.deleted_at, lce.deleted_at, lpe.deleted_at) AS deleted_at 
	FROM keypad_conditions k 
		LEFT JOIN keypad_conditions_to_lamp_events ktl on k.id = ktl.condition_id 
		LEFT JOIN 
			(SELECT 
				*, 
				'toggle' AS event_type 
			FROM lamp_toggle_events 
				WHERE deleted_at IS NULL) lte on ktl.event_id = lte.id 
		LEFT JOIN
			(SELECT 
				*, 
				'color' AS event_type 
			FROM lamp_color_events 
				WHERE deleted_at IS NULL) lce on ktl.event_id = lce.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'pulse' AS event_type 
			FROM lamp_pulse_events 
				WHERE deleted_at IS NULL) lpe on ktl.event_id = lpe.id 
	WHERE k.mac = ? and k.button_id = ?`

	rows, err := p.db.Raw(query, mac, buttonID).Rows()
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
