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
	rows, err := p.db.Table("keypad_conditions").Select("keypad_conditions.id, keypad_conditions.mac, keypad_conditions.button_id, keypad_conditions.created_at, keypad_conditions.updated_at, keypad_conditions.deleted_at, lamp_events.id, lamp_events.mac, lamp_events.event_type, lamp_events.red, lamp_events.green, lamp_events.blue, lamp_events.created_at, lamp_events.updated_at, lamp_events.deleted_at").Joins("left join conditions_to_events on keypad_conditions.id = conditions_to_events.condition_id").Joins("left join lamp_events on conditions_to_events.event_id = lamp_events.id").Where("conditions_to_events.interaction_id=?", id).Rows()

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
			&interactionDetails.LampEvent.Blue,
			&interactionDetails.LampEvent.CreatedAt,
			&interactionDetails.LampEvent.UpdatedAt,
			&interactionDetails.LampEvent.DeletedAt,
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
