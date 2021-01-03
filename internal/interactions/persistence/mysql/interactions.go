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
	rows, err := p.db.Raw("select k.id, k.mac, k.button_id, k.created_at, k.updated_at, k.deleted_at, coalesce(lte.event_type, lce.event_type, lpe.event_type) as event_type, IFNULL(lte.id,''), IFNULL(lte.mac,''), lte.created_at, lte.updated_at, lte.deleted_at, IFNULL(lce.id,''), IFNULL(lce.mac,''), IFNULL(lce.red,0), IFNULL(lce.green,0), IFNULL(lce.blue,0), lce.created_at, lce.updated_at, lce.deleted_at, IFNULL(lpe.id,''), IFNULL(lpe.mac,''), IFNULL(lpe.red,0), IFNULL(lpe.green,0), IFNULL(lpe.blue,0), lpe.created_at, lpe.updated_at, lpe.deleted_at from keypad_conditions_to_lamp_events ktl left join keypad_conditions k on ktl.condition_id = k.id left join (select *, 'toggle' as event_type from lamp_toggle_events where deleted_at is null) lte on ktl.event_id = lte.id left join (select *, 'color' as event_type from lamp_color_events where deleted_at is null) lce on ktl.event_id = lce.id left join (select *, 'pulse' as event_type from lamp_pulse_events where deleted_at is null) lpe on ktl.event_id = lpe.id where ktl.interaction_id = ?", id).Rows()

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
