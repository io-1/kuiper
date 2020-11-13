package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/io-1/kuiper/internal/interactions/persistence/mysql"
	"github.com/io-1/kuiper/internal/interactions/pubsub/mosquitto"
	"github.com/io-1/kuiper/internal/logger"
	"github.com/io-1/kuiper/internal/logger/logruslogger"
	"google.golang.org/grpc"

	interactions "github.com/io-1/kuiper/internal/interactions/servers"
	interactions_pb "github.com/io-1/kuiper/internal/pb/interactions"
)

var (
	Version     string
	Build       string
	showVersion *bool
	port        string
	log         logger.Logger
	ctx         context.Context
	server      *interactions.InteractionsServer
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		port = os.Getenv("PORT")
		dbConn := os.Getenv("DB_CONN")

		log = logruslogger.NewLogrusLogger(true)
		ctx = context.Background()
		persistence, err := mysql.NewMysqlPersistence(dbConn)
		if err != nil {
			log.Fatal(err)
		}

		server = interactions.NewInteractionsServer(persistence)
	}
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

	// FIXME: send alert message when keypad is pressed

	if *showVersion {
		fmt.Printf("settings server: version %s build %s", Version, Build)
	} else {
		pubSub := mosquitto.NewMosquittoPubSub(log)
		err := pubSub.NewKeypadListener(ctx, "keypad_listener", os.Getenv("KEYPAD_MQTT_URL"))
		if err != nil {
			log.Fatal(err)
		}

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Listening on port: %s\n", port)
		grpcServer := grpc.NewServer()
		interactions_pb.RegisterInteractionsServiceServer(grpcServer, server)
		grpcServer.Serve(lis)
	}
}
