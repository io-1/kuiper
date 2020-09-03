package mysql

import "github.com/n7down/kuiper/internal/devices/persistence"

func (p *MysqlPersistence) CreateBatCaveDeviceSetting(setting persistence.BatCaveDeviceSetting) int64 {
	rowsAffected := p.db.Create(&setting).RowsAffected
	return rowsAffected
}

func (p *MysqlPersistence) GetBatCaveDeviceSetting(deviceID string) (bool, persistence.BatCaveDeviceSetting) {
	var setting persistence.BatCaveDeviceSetting
	recordNotFound := p.db.Where("device_id=?", deviceID).First(&setting).RecordNotFound()
	return recordNotFound, setting
}

func (p *MysqlPersistence) UpdateBatCaveDeviceSetting(setting persistence.BatCaveDeviceSetting) int64 {
	rowsAffected := p.db.Model(&setting).Where("device_id = ?", setting.DeviceID).Updates(persistence.BatCaveDeviceSetting{DeepSleepDelay: setting.DeepSleepDelay}).RowsAffected
	return rowsAffected
}
