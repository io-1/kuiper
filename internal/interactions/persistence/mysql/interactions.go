package mysql

import "github.com/io-1/kuiper/internal/interactions/persistence"

func (p MysqlPersistence) CreateInteraction(interaction persistence.Interaction) int64 {
	rowsAffected := p.db.Create(&interaction).RowsAffected
	return rowsAffected
}

func (p MysqlPersistence) GetInteraction(id string) (recordNotFound bool, interaction persistence.Interaction) {
	recordNotFound = p.db.Where("id=?", id).First(&interaction).RecordNotFound()
	return recordNotFound, interaction
}

func (p MysqlPersistence) GetInteractionDetails(id string) ([]persistence.InteractionDetails, error) {
	query := `SELECT
		k.id, 
		k.mac, 
		k.button_id, 
		k.created_at, 
		k.updated_at, 
		k.deleted_at, 
		COALESCE(lte.id, lce.id, lpe.id) AS id, 
		COALESCE(lte.mac, lce.mac, lpe.mac) AS mac, 
		COALESCE(lte.event_type, lce.event_type, lpe.event_type) AS event_type, 
		IFNULL(COALESCE(lce.red, lpe.red), 0), 
		IFNULL(COALESCE(lce.green, lpe.green), 0), 
		IFNULL(COALESCE(lce.blue, lpe.blue), 0), 
		COALESCE(lte.created_at, lce.created_at, lpe.created_at) AS created_at, 
		COALESCE(lte.updated_at, lce.updated_at, lpe.updated_at) AS updated_at, 
		COALESCE(lte.deleted_at, lce.deleted_at, lpe.deleted_at) AS deleted_at 
	FROM keypad_conditions_to_lamp_events ktl 
		LEFT JOIN keypad_conditions k ON ktl.condition_id = k.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'toggle' AS event_type 
			FROM lamp_toggle_events WHERE deleted_at is null) lte ON ktl.event_id = lte.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'color' AS event_type 
			FROM lamp_color_events WHERE deleted_at IS null) lce ON ktl.event_id = lce.id 
		LEFT JOIN 
			(SELECT 
				*, 
				'pulse' AS event_type
			FROM lamp_pulse_events WHERE deleted_at IS null) lpe ON ktl.event_id = lpe.id 
	WHERE ktl.interaction_id = ?`

	rows, err := p.db.Raw(query, id).Rows()

	defer rows.Close()
	if err != nil {
		return []persistence.InteractionDetails{}, err
	}

	var allInteractionDetails []persistence.InteractionDetails
	for rows.Next() {
		var interactionDetails persistence.InteractionDetails
		err = rows.Scan(
			&interactionDetails.KeypadCondition.ID,
			&interactionDetails.KeypadCondition.Mac,
			&interactionDetails.KeypadCondition.ButtonID,
			&interactionDetails.KeypadCondition.CreatedAt,
			&interactionDetails.KeypadCondition.UpdatedAt,
			&interactionDetails.KeypadCondition.DeletedAt,
			&interactionDetails.LampEvent.ID,
			&interactionDetails.LampEvent.Mac,
			&interactionDetails.LampEvent.EventType,
			&interactionDetails.LampEvent.Red,
			&interactionDetails.LampEvent.Green,
			&interactionDetails.LampEvent.Blue, &interactionDetails.LampEvent.CreatedAt, &interactionDetails.LampEvent.UpdatedAt, &interactionDetails.LampEvent.DeletedAt,
		)

		if err != nil {
			return []persistence.InteractionDetails{}, err
		}

		allInteractionDetails = append(allInteractionDetails, interactionDetails)
	}

	return allInteractionDetails, nil
}

func (p MysqlPersistence) UpdateInteraction(interaction persistence.Interaction) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", interaction.ID).First(&persistence.Interaction{}).RecordNotFound()
	err = p.db.Model(&interaction).Where("id=?", interaction.ID).Updates(persistence.Interaction{Name: interaction.Name, Description: interaction.Description}).Error
	return recordNotFound, err
}

func (p MysqlPersistence) DeleteInteraction(interaction persistence.Interaction) (recordNotFound bool, err error) {
	recordNotFound = p.db.Where("id=?", interaction.ID).First(&persistence.Interaction{}).RecordNotFound()
	err = p.db.Delete(&interaction).Error
	return recordNotFound, err
}
