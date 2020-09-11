package request

import (
	"github.com/io-1/kuiper/internal/devices/commands"
	"github.com/io-1/kuiper/internal/devices/persistence"
	"github.com/io-1/kuiper/internal/devices/pubsub/response"
)

type BatCaveDeviceSettingRequest struct {
	Mac            string `json:"m"`
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
