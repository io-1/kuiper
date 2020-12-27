package mosquitto

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/io-1/kuiper/internal/logger"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

type MosquittoPublisher struct {
	logger logger.Logger
	client mqtt.Client
}

func NewMosquittoPublisher(logger logger.Logger, mqttURL, publisherName string) (*MosquittoPublisher, error) {
	mqttUrl, err := url.Parse(mqttURL)
	if err != nil {
		return &MosquittoPublisher{}, err
	}

	// topic := mqttUrl.Path[1:len(mqttUrl.Path)]
	// if topic == "" {
	// 	topic = "test"
	// }

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", mqttUrl.Host))
	opts.SetUsername(mqttUrl.User.Username())
	password, _ := mqttUrl.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(publisherName)

	// var f mqtt.MessageHandler = p.BatCaveDeviceSettingsListenerMessageHandler

	// opts.SetDefaultPublishHandler(f)

	// err = nil
	// opts.OnConnect = func(c mqtt.Client) {
	// 	if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
	// 		err = token.Error()
	// 	}
	// }

	// if err != nil {
	// 	return err
	// }

	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return &MosquittoPublisher{}, errors.New(fmt.Sprintf("Error with %s: %s", publisherName, token.Error()))
	}
	fmt.Println(fmt.Sprintf("%s on %s is connected", publisherName, mqttUrl.String()))
	return &MosquittoPublisher{
		logger: logger,
		client: mqttClient,
	}, nil
}

func (p MosquittoPublisher) Publish(topic string, json []byte) {
	p.logger.Infof("Sending message %s to %s", json, topic)
	token := p.client.Publish(topic, 0, false, json)
	token.WaitTimeout(ONE_MINUTE)
}
