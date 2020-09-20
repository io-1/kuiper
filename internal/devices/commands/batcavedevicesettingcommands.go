package commands

import (
	"fmt"
	"strconv"
)

const (
	DEEP_SLEEP_DELAY_COMMAND = 0x00
)

type BatCaveDeviceSettingCommands struct {
	commands []string
}

func (c *BatCaveDeviceSettingCommands) GetCommands() []string {
	return c.commands
}

func (c *BatCaveDeviceSettingCommands) GetCommandsInt() []int {
	commandsInt := []int{}
	for _, command := range c.commands {
		commandInt, _ := strconv.ParseInt(command, 16, 64)
		commandsInt = append(commandsInt, int(commandInt))
	}
	return commandsInt
}

func (c *BatCaveDeviceSettingCommands) AddDeepSleepDelayCommand(d uint32) {
	h := fmt.Sprintf("%04x%04x", DEEP_SLEEP_DELAY_COMMAND, d)
	c.commands = append(c.commands, h)
}
