package main

import (
	"context"
	"os"

	"github.com/io-1/kuiper/internal/interactions/pubsub/mosquitto"
	"github.com/io-1/kuiper/internal/logger/logruslogger"
)

func init() {
}

func main() {

	// FIXME: keep one record of all sensor data
	// FIXME: keeps track of all the interactions that need to happen
	// FIXME: check after a sensor is updated to any interactions on involved with the device
	// FIXME: send message to devices if interactions are met
	// FIXME: implement scheduled events and one-time events
	// where scheduled events happen on a schedule - ie monday at 5 pm
	// one-time happen one time - ie turn off device for the rest of the day

	// FIXME: keypad - when a button press comes through - dont store it
	// FIXME: just check if there are any interactions tied to it - and execute them

	// FIXME: create keypad listener
	log := logruslogger.NewLogrusLogger(true)
	ctx := context.Background()
	pubSub := mosquitto.NewMosquittoPubSub(log)
	err := pubSub.NewKeypadListener(ctx, "keypad_listener", os.Getenv("KEYPAD_MQTT_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// FIXME: send alert message when keypad is pressed
}
