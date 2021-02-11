package mosquitto

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	lamp_events "github.com/io-1/kuiper/internal/events/lamp"
	sensors "github.com/io-1/kuiper/internal/interactions/devicesensors"
)

const (
	ONE_MINUTE                        = 1 * time.Minute
	LAMP_ON_EVENT                     = "on"
	LAMP_OFF_EVENT                    = "off"
	LAMP_TOGGLE_EVENT                 = "toggle"
	LAMP_BRIGHTNESS_EVENT             = "brightness"
	LAMP_AUTO_BRIGHTNESS_ON_EVENT     = "auto-brightness-on"
	LAMP_AUTO_BRIGHTNESS_OFF_EVENT    = "auto-brightness-off"
	LAMP_AUTO_BRIGHTNESS_TOGGLE_EVENT = "auto-brightness-toggle"
	LAMP_COLOR_EVENT                  = "color"
	LAMP_PULSE_EVENT                  = "pulse"
)

func (p MosquittoPubSub) NewKeypadListener(ctx context.Context, listenerName string, subscription string) error {
	mqttUrl, err := url.Parse(subscription)
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

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		p.logger.Infof("Received message: %s\n", msg.Payload())

		// unmashal payload
		sensor := &sensors.KeypadMeasurement{}
		err := json.Unmarshal([]byte(msg.Payload()), sensor)
		if err != nil {
			p.logger.Error(err.Error())
			return
		}

		p.logger.Infof("Unmashalled message: %v\n", sensor)

		// FIXME: check the cache - and return it if the data is cached

		lampEvents, err := p.persistence.GetLampEventsByKeypadMacAndButtonID(sensor.Mac, sensor.ID)
		if err != nil {
			p.logger.Error(err.Error())
			return
		}

		// FIXME: cache the data

		// for each lamp event - send event to the device
		for _, lampEvent := range lampEvents {
			var lampEventToSend interface{}
			switch lampEvent.EventType {
			case LAMP_ON_EVENT:
				lampEventToSend = lamp_events.NewLampDeviceOnEvent()
			case LAMP_OFF_EVENT:
			case LAMP_TOGGLE_EVENT:
				lampEventToSend = lamp_events.NewLampDeviceToggleEvent()
			case LAMP_BRIGHTNESS_EVENT:
				lampEventToSend = lamp_events.NewLampDeviceBrightnessEvent(lampEvent.Brightness)
			case LAMP_AUTO_BRIGHTNESS_ON_EVENT:
				lampEventToSend = lamp_events.NewLampDeviceAutoBrightnessOnEvent()
			case LAMP_AUTO_BRIGHTNESS_OFF_EVENT:
				lampEventToSend = lamp_events.NewLampDeviceAutoBrightnessOffEvent()
			case LAMP_AUTO_BRIGHTNESS_TOGGLE_EVENT:
				lampEventToSend = lamp_events.NewLampDeviceAutoBrightnessToggleEvent()
			case LAMP_COLOR_EVENT:
				lampEventToSend = lamp_events.NewLampDeviceColorEvent(lampEvent.Red, lampEvent.Green, lampEvent.Blue)
			case LAMP_PULSE_EVENT:
				lampEventToSend = lamp_events.NewLampDevicePulseEvent(lampEvent.Red, lampEvent.Green, lampEvent.Blue)
			}

			json, err := json.Marshal(lampEventToSend)
			if err != nil {
				p.logger.Error(err)
				return
			}

			deviceTopic := fmt.Sprintf("devices/%s", lampEvent.Mac)
			p.logger.Infof("Sending message %s to %s", json, deviceTopic)
			token := client.Publish(deviceTopic, 0, false, json)
			token.WaitTimeout(ONE_MINUTE)
		}
	}

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
