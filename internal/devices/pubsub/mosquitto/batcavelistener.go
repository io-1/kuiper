package mosquitto

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/io-1/kuiper/internal/devices/listeners/request"
	"github.com/io-1/kuiper/internal/devices/listeners/response"
	"github.com/io-1/kuiper/internal/devices/persistence"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

func (p MosquittoPubSub) BatCaveDeviceSettingsListenerMessageHandler(client mqtt.Client, msg mqtt.Message) {
	p.logger.Infof("Received message: %s\n", msg.Payload())

	// unmashal payload
	var (
		req request.BatCaveDeviceSettingRequest
		res response.BatCaveDeviceSettingResponse
	)

	err := json.Unmarshal([]byte(msg.Payload()), &req)
	if err != nil {
		p.logger.Error(err)
		return
	}

	// get the settings
	recordNotFound, settingInPersistence := p.persistence.GetBatCaveDeviceSettingByMac(req.Mac)
	if recordNotFound {

		// send back default values
		res = response.GetBatCaveDeviceSettingDefault()

		newSetting := persistence.BatCaveDeviceSetting{
			ID:             uuid.New().String(),
			Mac:            req.Mac,
			DeepSleepDelay: res.DeepSleepDelay,
		}

		// create the new setting
		p.persistence.CreateBatCaveDeviceSetting(newSetting)

	} else {

		// check for the differences in the settings
		var isEqual bool
		isEqual, res = req.IsEqual(settingInPersistence)
		p.logger.Infof("Settings are equal: %t - %v %v", isEqual, settingInPersistence, res)
		if isEqual {

			// settings are the same on the device and in persistence - return
			return
		}
	}

	json, err := json.Marshal(res)
	if err != nil {
		p.logger.Error(err)
		return
	}

	// send back to the device the new settings
	deviceTopic := fmt.Sprintf("devices/%s", req.Mac)
	p.logger.Infof("Sending message %s to %s", json, deviceTopic)
	token := client.Publish(deviceTopic, 0, false, json)
	token.WaitTimeout(ONE_MINUTE)
}

func (p MosquittoPubSub) NewBatCaveDeviceSettingsListener(listenerName string, mqttURL string) error {
	mqttUrl, err := url.Parse(mqttURL)
	if err != nil {
		return err
	}

	topic := mqttUrl.Path[1:len(mqttUrl.Path)]
	if topic == "" {
		topic = "test"
	}

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", mqttUrl.Host))
	opts.SetUsername(mqttUrl.User.Username())
	password, _ := mqttUrl.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(listenerName)

	var f mqtt.MessageHandler = p.BatCaveDeviceSettingsListenerMessageHandler

	opts.SetDefaultPublishHandler(f)

	err = nil
	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			err = token.Error()
		}
	}

	if err != nil {
		return err
	}

	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return errors.New(fmt.Sprintf("Error with %s: %s", listenerName, token.Error()))
	}
	fmt.Println(fmt.Sprintf("%s on %s is connected", listenerName, mqttUrl.String()))

	return nil
}
