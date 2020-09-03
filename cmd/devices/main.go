package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/n7down/kuiper/internal/devices/persistence/mysql"
	"github.com/n7down/kuiper/internal/devices/pubsub/mosquitto"
	"github.com/n7down/kuiper/internal/logger"
	"github.com/n7down/kuiper/internal/logger/logruslogger"
	"google.golang.org/grpc"

	devices "github.com/n7down/kuiper/internal/devices/servers"
	devices_pb "github.com/n7down/kuiper/internal/pb/devices"
)

const (
	ONE_MINUTE = 1 * time.Minute
)

var (
	Version     string
	Build       string
	showVersion *bool
	port        string
	log         logger.Logger
	server      *devices.DevicesServer
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		port = os.Getenv("PORT")
		dbConn := os.Getenv("DB_CONN")
		batCaveSettingsMQTTURL := os.Getenv("BAT_CAVE_SETTINGS_MQTT_URL")

		log = logruslogger.NewLogrusLogger(true)
		persistence, err := mysql.NewMysqlPersistence(dbConn)
		if err != nil {
			log.Fatal(err)
		}

		server = devices.NewDevicesServer(persistence)
		pubSub := mosquitto.NewMosquittoPubSub(persistence, log)
		err = pubSub.NewBatCaveDeviceSettingsListener("bat_cave_settings_listener", batCaveSettingsMQTTURL)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	if *showVersion {
		fmt.Printf("settings server: version %s build %s", Version, Build)
	} else {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Listening on port: %s\n", port)
		grpcServer := grpc.NewServer()
		devices_pb.RegisterDevicesServiceServer(grpcServer, server)
		grpcServer.Serve(lis)
	}
}
