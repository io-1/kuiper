package listeners

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	listeners "github.com/io-1/kuiper/internal/common/listeners"
	sensors "github.com/io-1/kuiper/internal/sensors/devicesensors"
)

func (e SensorsListenersEnv) NewKeypadListener(listenerName string, urlString string) (*listeners.Listener, error) {
	i := &listeners.Listener{}

	mqttUrl, err := url.Parse(urlString)
	if err != nil {
		return i, err
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
		logrus.Infof("Received message: %s\n", msg.Payload())

		// unmashal payload
		measurement := &sensors.KeypadMeasurement{}
		err := json.Unmarshal([]byte(msg.Payload()), measurement)
		if err != nil {
			logrus.Error(err.Error())
		}

		if err == nil {
			err = e.persistence.CreateButtonMeasurement(measurement)
			logrus.Infof("Logged measurement: %v", measurement)
			if err != nil {
				logrus.Error(err.Error())
			}
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
		return i, err
	}

	i.MqttOptions = opts
	i.ListenerName = listenerName
	i.ListenerMQTTUrl = mqttUrl

	return i, nil
}
