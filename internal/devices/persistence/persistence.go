//go:generate mockgen -source persistence.go -destination=mock/mockpersistence.go -package=mock
package persistence

type Persistence interface {
	CreateBatCaveDeviceSetting(settings BatCaveDeviceSetting) int64
	GetBatCaveDeviceSetting(deviceID string) (bool, BatCaveDeviceSetting)
	UpdateBatCaveDeviceSetting(settings BatCaveDeviceSetting) int64
}
