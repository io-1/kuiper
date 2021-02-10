package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername("mqtt")
	password := ""
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

type WeatherMeasurement struct {
	Sensor string `json:"sensor"`
	Mac    string `json:"mac"`
	Temp   int    `json:"temperature"`
}

func main() {
	uri, err := url.Parse("tcp://localhost:1883")
	if err != nil {
		log.Fatal(err)
	}

	client := connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)

	for t := range timer.C {
		var (
			min    int = 10
			max    int = 100
			random int = rand.Intn(max-min) + min
			mac        = "ff0022334455"
		)

		w := WeatherMeasurement{
			Sensor: "weather",
			Mac:    mac,
			Temp:   random,
		}

		// nsec := time.Now().UnixNano()
		// payload := "weather,location=us-midwest temperature=" + strconv.FormatInt(random, 10) + " " + strconv.FormatInt(nsec, 10)
		// payload := fmt.Sprintf("weather,location=us-midwest temperature=%d", random)

		j, err := json.Marshal(&w)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fmt.Sprintf("time: %v payload: %v", t, w))
		client.Publish("sensors/weather", 0, false, j)
		// fmt.Println(fmt.Sprintf("time: %v payload: %v", t, payload))
		// client.Publish("sensors", 0, false, payload)
	}
}
