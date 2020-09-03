package mysql

import "github.com/n7down/kuiper/internal/devices/persistence"

func (p *MysqlPersistence) CreateBatCaveSetting(setting persistence.BatCaveSetting) int64 {
	rowsAffected := p.db.Create(&setting).RowsAffected
	return rowsAffected
}

func (p *MysqlPersistence) GetBatCaveSetting(deviceID string) (bool, persistence.BatCaveSetting) {
	var setting persistence.BatCaveSetting
	recordNotFound := p.db.Where("device_id=?", deviceID).First(&setting).RecordNotFound()
	return recordNotFound, setting
}

func (p *MysqlPersistence) UpdateBatCaveSetting(setting persistence.BatCaveSetting) int64 {
	rowsAffected := p.db.Model(&setting).Where("device_id = ?", setting.DeviceID).Updates(persistence.BatCaveSetting{DeepSleepDelay: setting.DeepSleepDelay}).RowsAffected
	return rowsAffected
}
