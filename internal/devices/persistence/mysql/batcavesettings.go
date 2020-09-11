package mysql

import "github.com/io-1/kuiper/internal/devices/persistence"

func (p *MysqlPersistence) CreateBatCaveDeviceSetting(setting persistence.BatCaveDeviceSetting) int64 {
	rowsAffected := p.db.Create(&setting).RowsAffected
	return rowsAffected
}

func (p *MysqlPersistence) GetBatCaveDeviceSetting(id string) (bool, persistence.BatCaveDeviceSetting) {
	var setting persistence.BatCaveDeviceSetting
	recordNotFound := p.db.Where("id=?", id).First(&setting).RecordNotFound()
	return recordNotFound, setting
}

func (p *MysqlPersistence) GetBatCaveDeviceSettingByMac(mac string) (bool, persistence.BatCaveDeviceSetting) {
	var setting persistence.BatCaveDeviceSetting
	recordNotFound := p.db.Where("mac=?", mac).First(&setting).RecordNotFound()
	return recordNotFound, setting
}

func (p *MysqlPersistence) UpdateBatCaveDeviceSetting(setting persistence.BatCaveDeviceSetting) int64 {
	rowsAffected := p.db.Model(&setting).Where("id=?", setting.ID).Updates(persistence.BatCaveDeviceSetting{Mac: setting.Mac, DeepSleepDelay: setting.DeepSleepDelay}).RowsAffected
	return rowsAffected
}
