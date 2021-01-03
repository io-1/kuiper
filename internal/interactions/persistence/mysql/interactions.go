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
	// rows, err := p.db.Table("keypad_conditions").Select("keypad_conditions.id, keypad_conditions.mac, keypad_conditions.button_id, keypad_conditions.created_at, keypad_conditions.updated_at, keypad_conditions.deleted_at, lamp_events.id, lamp_events.mac, lamp_events.event_type, lamp_events.red, lamp_events.green, lamp_events.blue, lamp_events.created_at, lamp_events.updated_at, lamp_events.deleted_at").Joins("left join conditions_to_events on keypad_conditions.id = conditions_to_events.condition_id").Joins("left join lamp_events on conditions_to_events.event_id = lamp_events.id").Where("conditions_to_events.interaction_id=?", id).Rows()

	// FIXME: unable to scan a date instead of a time
	rows, err := p.db.Raw("select k.id, k.mac, k.button_id, k.created_at, k.updated_at, k.deleted_at, coalesce(lte.event_type, lce.event_type, lpe.event_type) as event_type, IFNULL(lte.id,''), IFNULL(lte.mac,''), IFNULL(lte.created_at,'1970-12-01'), IFNULL(lte.updated_at,'1970-12-01'), IFNULL(lte.deleted_at,'1970-12-01'), IFNULL(lce.id,''), IFNULL(lce.mac,''), IFNULL(lce.red,0), IFNULL(lce.green,0), IFNULL(lce.blue,0), IFNULL(lce.created_at,'1970-12-01'), IFNULL(lce.updated_at,'1970-12-01'), IFNULL(lce.deleted_at,'1970-12-01'), IFNULL(lpe.id,''), IFNULL(lpe.mac,''), IFNULL(lpe.red,0), IFNULL(lpe.green,0), IFNULL(lpe.blue,0), IFNULL(lpe.created_at,'1970-12-01'), IFNULL(lpe.updated_at,'1970-12-01'), IFNULL(lpe.deleted_at,'1970-12-01') from keypad_conditions_to_lamp_events ktl left join keypad_conditions k on ktl.condition_id = k.id left join (select *, 'toggle' as event_type from lamp_toggle_events where deleted_at is null) lte on ktl.event_id = lte.id left join (select *, 'color' as event_type from lamp_color_events where deleted_at is null) lce on ktl.event_id = lce.id left join (select *, 'pulse' as event_type from lamp_pulse_events where deleted_at is null) lpe on ktl.event_id = lpe.id where ktl.interaction_id = ?", id).Rows()

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

			&interactionDetails.LampEventType,

			&interactionDetails.LampToggleEvent.ID,
			&interactionDetails.LampToggleEvent.Mac,
			&interactionDetails.LampToggleEvent.CreatedAt,
			&interactionDetails.LampToggleEvent.UpdatedAt,
			&interactionDetails.LampToggleEvent.DeletedAt,

			&interactionDetails.LampColorEvent.ID,
			&interactionDetails.LampColorEvent.Mac,
			&interactionDetails.LampColorEvent.Red,
			&interactionDetails.LampColorEvent.Green,
			&interactionDetails.LampColorEvent.Blue,
			&interactionDetails.LampColorEvent.CreatedAt,
			&interactionDetails.LampColorEvent.UpdatedAt,
			&interactionDetails.LampColorEvent.DeletedAt,

			&interactionDetails.LampPulseEvent.ID,
			&interactionDetails.LampPulseEvent.Mac,
			&interactionDetails.LampPulseEvent.Red,
			&interactionDetails.LampPulseEvent.Green,
			&interactionDetails.LampPulseEvent.Blue,
			&interactionDetails.LampPulseEvent.CreatedAt,
			&interactionDetails.LampPulseEvent.UpdatedAt,
			&interactionDetails.LampPulseEvent.DeletedAt,
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
