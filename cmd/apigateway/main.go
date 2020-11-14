package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/io-1/kuiper/internal/apigateway"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions"
	"github.com/io-1/kuiper/internal/apigateway/serviceinfo"
	"github.com/io-1/kuiper/internal/logger"
	"github.com/io-1/kuiper/internal/logger/logruslogger"

	ginauth "github.com/io-1/kuiper/internal/apigateway/auth/ginauth"
	devices "github.com/io-1/kuiper/internal/apigateway/clients/devices"
	users "github.com/io-1/kuiper/internal/apigateway/clients/users"
)

var (
	Version string
	Build   string

	showVersion *bool
	log         logger.Logger

	env              string
	port             string
	devicesHost      string
	usersHost        string
	interactionsHost string
)

func init() {
	showVersion = flag.Bool("v", false, "show version and build")
	flag.Parse()
	if !*showVersion {
		log = logruslogger.NewLogrusLogger(true)
		env = os.Getenv("ENV")
		port = os.Getenv("PORT")
		devicesHost = os.Getenv("DEVICES_HOST")
		usersHost = os.Getenv("USERS_HOST")
		interactionsHost = os.Getenv("INTERACTIONS_HOST")
	}
}

func main() {
	if *showVersion {
		serviceInfo := serviceinfo.ServiceInfo{
			Version: Version,
			Build:   Build,
		}

		serviceInfoJson, err := json.Marshal(serviceInfo)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(string(serviceInfoJson))
	} else {

		devicesClient, err := devices.NewDevicesClient(devicesHost, log)
		if err != nil {
			log.Fatal(err)
		}

		usersClient, err := users.NewUsersClient(usersHost, log)
		if err != nil {
			log.Fatal(err)
		}

		interactionsClient, err := interactions.NewInteractionsClient(interactionsHost, log)
		if err != nil {
			log.Fatal(err)
		}

		ginAuth, err := ginauth.NewGinAuth(usersClient, log)
		if err != nil {
			log.Fatal(err)
		}

		apiGateway := apigateway.NewAPIGateway(
			env,
			Version,
			Build,
			ginAuth,
			devicesClient,
			usersClient,
			interactionsClient,
		)

		router := gin.Default()

		err = apiGateway.InitV1Routes(router)
		if err != nil {
			log.Fatal(err)
		}

		routerPort := fmt.Sprintf(":%s", port)
		log.Infof("Listening on port: %s\n", port)
		err = apiGateway.Run(router, routerPort)
		if err != nil {
			log.Fatal(err)
		}
	}
}
