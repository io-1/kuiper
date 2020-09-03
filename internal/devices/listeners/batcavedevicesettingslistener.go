package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/n7down/kuiper/internal/devices/listeners/request"
	"github.com/n7down/kuiper/internal/devices/listeners/response"
	"github.com/n7down/kuiper/internal/devices/persistence"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	listeners "github.com/n7down/kuiper/internal/common/listeners"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

func (e *DeviceSettingsListenersEnv) BatCaveDeviceSettingsListenerMessageHandler(client mqtt.Client, msg mqtt.Message) {
	e.logger.Infof("Received message: %s\n", msg.Payload())

	// unmashal payload
	var (
		req request.BatCaveDeviceSettingRequest
		res response.BatCaveDeviceSettingResponse
	)

	err := json.Unmarshal([]byte(msg.Payload()), &req)
	if err != nil {
		e.logger.Error(err)
		return
	}

	// get the settings
	recordNotFound, settingInPersistence := e.persistence.GetBatCaveDeviceSetting(req.DeviceID)
	if recordNotFound {

		// send back default values
		res = response.GetBatCaveDeviceSettingDefault()

		newSetting := persistence.BatCaveDeviceSetting{
			DeviceID:       req.DeviceID,
			DeepSleepDelay: res.DeepSleepDelay,
		}

		// create the new setting
		e.persistence.CreateBatCaveDeviceSetting(newSetting)

	} else {

		// check for the differences in the settings
		var isEqual bool
		isEqual, res = req.IsEqual(settingInPersistence)
		e.logger.Infof("Settings are equal: %t - %v %v", isEqual, settingInPersistence, res)
		if isEqual {

			// settings are the same on the device and in persistence - return
			return
		}
	}

	json, err := json.Marshal(res)
	if err != nil {
		e.logger.Error(err)
		return
	}

	// send back to the device the new settings
	deviceTopic := fmt.Sprintf("devices/%s", req.DeviceID)
	e.logger.Infof("Sending message %s to %s", json, deviceTopic)
	token := client.Publish(deviceTopic, 0, false, json)
	token.WaitTimeout(ONE_MINUTE)
}

func (e *DeviceSettingsListenersEnv) NewBatCaveDeviceSettingsListener(listenerName string, mqttURL string) (*listeners.Listener, error) {
	i := &listeners.Listener{}

	u, err := url.Parse(mqttURL)
	if err != nil {
		return i, err
	}

	topic := u.Path[1:len(u.Path)]
	if topic == "" {
		topic = "test"
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", u.Host))
	opts.SetUsername(u.User.Username())
	password, _ := u.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(listenerName)

	var f mqtt.MessageHandler = e.BatCaveDeviceSettingsListenerMessageHandler

	opts.SetDefaultPublishHandler(f)

	err = nil
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			err = token.Error()
		}
	}

	if err != nil {
		return i, err
	}

	i.MqttOptions = opts
	i.ListenerName = listenerName
	i.ListenerMQTTUrl = u

	return i, nil
}
