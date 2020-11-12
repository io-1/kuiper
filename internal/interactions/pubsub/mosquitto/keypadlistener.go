package mosquitto

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	sensors "github.com/io-1/kuiper/internal/interactions/devicesensors"
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
		}

		// if err == nil {
		// 	err = p.persistence.CreateButtonMeasurement(sensor)
		// 	p.logger.Infof("Logged sensor: %v", sensor)
		// 	if err != nil {
		// 		p.logger.Error(err.Error())
		// 	}
		// }

		// FIXME: when a button interaction comes in check the database for interactions
		// keypadConditions, err = p.persistence.GetKeypadConditions(sensor.ID, sensor.Mac)
		// if err != nil {
		// 	p.logger.Error(err.Error())
		// }

		// FIXME: send off the button interactoins by the mac
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
