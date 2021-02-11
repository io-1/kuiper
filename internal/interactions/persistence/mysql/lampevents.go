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
		COALESCE(loe.id, lfe.id, lte.id, lbe.id, labone.id, laboffe.id, labte.id, lce.id, lpe.id) AS id, 
		COALESCE(loe.mac, lfe.mac, lte.mac, lbe.mac, labone.mac, laboffe.mac, labte.mac, lce.mac, lpe.mac) AS mac, 
		COALESCE(loe.event_type, lfe.event_type, lte.event_type, lbe.event_type, labone.event_type, laboffe.event_type, labte.event_type, lce.event_type, lpe.event_type) AS event_type, 
		IFNULL(COALESCE(lce.red, lpe.red),0) AS red, 
		IFNULL(COALESCE(lce.green, lpe.green),0) AS green, 
		IFNULL(COALESCE(lce.blue, lpe.blue),0) AS blue,
		IFNULL(lbe.brightness,0) AS brightness,
		COALESCE(loe.created_at, lfe.created_at, lte.created_at, labone.created_at, laboffe.created_at, labte.created_at, lce.created_at, lpe.created_at) AS created_at, 
		COALESCE(loe.updated_at, lfe.updated_at, lte.updated_at, labone.updated_at, laboffe.updated_at, labte.updated_at, lce.updated_at, lpe.updated_at) AS updated_at, 
		COALESCE(loe.deleted_at, lfe.deleted_at, lte.deleted_at, labone.deleted_at, laboffe.deleted_at, labte.deleted_at, lce.deleted_at, lpe.deleted_at) AS deleted_at 
	FROM keypad_conditions_to_lamp_events ktl 
		LEFT JOIN 
			(SELECT 
				*, 
				'on' AS event_type 
			FROM lamp_on_events 
				WHERE deleted_at IS NULL) loe on ktl.event_id = loe.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'off' AS event_type 
			FROM lamp_off_events 
				WHERE deleted_at IS NULL) lfe on ktl.event_id = lfe.id 
		left join 
			(SELECT 
				*, 
				'toggle' AS event_type 
			FROM lamp_toggle_events WHERE deleted_at IS NULL) lte ON ktl.event_id = lte.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'brightness' AS event_type 
			FROM lamp_brightness_events 
				WHERE deleted_at IS NULL) lbe on ktl.event_id = lbe.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'auto-brightness-on' AS event_type 
			FROM lamp_auto_brightness_on_events 
				WHERE deleted_at IS NULL) labone on ktl.event_id = labone.id 

		LEFT JOIN 
			(SELECT 
				*, 
				'auto-brightness-off' AS event_type 
			FROM lamp_auto_brightness_off_events 
				WHERE deleted_at IS NULL) laboffe on ktl.event_id = laboffe.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'auto-brightness-toggle' AS event_type 
			FROM lamp_auto_brightness_toggle_events 
				WHERE deleted_at IS NULL) labte on ktl.event_id = labte.id 
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
		COALESCE(loe.id, lfe.id, lte.id, lbe.id, labone.id, laboffe.id, labte.id, lce.id, lpe.id) AS id, 
		COALESCE(loe.mac, lfe.mac, lte.mac, lbe.mac, labone.mac, laboffe.mac, labte.mac, lce.mac, lpe.mac) AS mac, 
		COALESCE(loe.event_type, lfe.event_type, lte.event_type, lbe.event_type, labone.event_type, laboffe.event_type, labte.event_type, lce.event_type, lpe.event_type) AS event_type, 
		IFNULL(COALESCE(lce.red, lpe.red),0) AS red, 
		IFNULL(COALESCE(lce.green, lpe.green),0) AS green, 
		IFNULL(COALESCE(lce.blue, lpe.blue),0) AS blue,
		IFNULL(lbe.brightness,0) AS brightness,
		COALESCE(loe.created_at, lfe.created_at, lte.created_at, labone.created_at, laboffe.created_at, labte.created_at, lce.created_at, lpe.created_at) AS created_at, 
		COALESCE(loe.updated_at, lfe.updated_at, lte.updated_at, labone.updated_at, laboffe.updated_at, labte.updated_at, lce.updated_at, lpe.updated_at) AS updated_at, 
		COALESCE(loe.deleted_at, lfe.deleted_at, lte.deleted_at, labone.deleted_at, laboffe.deleted_at, labte.deleted_at, lce.deleted_at, lpe.deleted_at) AS deleted_at 
	FROM keypad_conditions k 
		LEFT JOIN keypad_conditions_to_lamp_events ktl on k.id = ktl.condition_id 
		LEFT JOIN 
			(SELECT 
				*, 
				'on' AS event_type 
			FROM lamp_on_events 
				WHERE deleted_at IS NULL) loe on ktl.event_id = loe.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'off' AS event_type 
			FROM lamp_off_events 
				WHERE deleted_at IS NULL) lfe on ktl.event_id = lfe.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'toggle' AS event_type 
			FROM lamp_toggle_events 
				WHERE deleted_at IS NULL) lte on ktl.event_id = lte.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'brightness' AS event_type 
			FROM lamp_brightness_events 
				WHERE deleted_at IS NULL) lbe on ktl.event_id = lbe.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'auto-brightness-on' AS event_type 
			FROM lamp_auto_brightness_on_events 
				WHERE deleted_at IS NULL) labone on ktl.event_id = labone.id 

		LEFT JOIN 
			(SELECT 
				*, 
				'auto-brightness-off' AS event_type 
			FROM lamp_auto_brightness_off_events 
				WHERE deleted_at IS NULL) laboffe on ktl.event_id = laboffe.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'auto-brightness-toggle' AS event_type 
			FROM lamp_auto_brightness_toggle_events 
				WHERE deleted_at IS NULL) labte on ktl.event_id = labte.id 
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
			&lampEvent.Brightness,
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
