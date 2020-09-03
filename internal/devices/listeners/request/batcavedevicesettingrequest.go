package request

import (
	"github.com/n7down/kuiper/internal/devices/listeners/commands"
	"github.com/n7down/kuiper/internal/devices/listeners/response"
	"github.com/n7down/kuiper/internal/devices/persistence"
)

type BatCaveDeviceSettingRequest struct {
	DeviceID       string `json:"m"`
	DeepSleepDelay uint32 `json:"s"`
}

func (s *BatCaveDeviceSettingRequest) IsEqual(settings persistence.BatCaveDeviceSetting) (bool, response.BatCaveDeviceSettingResponse) {
	res := response.BatCaveDeviceSettingResponse{}
	isEqual := true

	if s.DeepSleepDelay != settings.DeepSleepDelay {
		isEqual = false
		res.DeepSleepDelay = settings.DeepSleepDelay
	} else {
		res.DeepSleepDelay = s.DeepSleepDelay
	}

	return isEqual, res
}

func (s *BatCaveDeviceSettingRequest) IsEqualAndGetCommands(settings persistence.BatCaveDeviceSetting) (bool, []string) {
	commands := commands.BatCaveDeviceSettingCommands{}
	hasChanges := false

	if s.DeepSleepDelay != settings.DeepSleepDelay {
		hasChanges = true
		commands.AddDeepSleepDelayCommand(settings.DeepSleepDelay)
	}

	return hasChanges, commands.GetCommands()
}
