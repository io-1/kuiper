package lamp

// FIXME: change from response to events
// FIXME: move to /internal/events/lamp dir
// FIXME: these events will be used by the devices and interaction service
type LampDeviceAutoBrightnessEvent struct {
	EventType string `json:"e"`
}
