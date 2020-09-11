package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway"

	ginauth "github.com/io-1/kuiper/internal/apigateway/auth/ginauth"
	devices "github.com/io-1/kuiper/internal/apigateway/clients/devices"
	users "github.com/io-1/kuiper/internal/apigateway/clients/users"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	port := os.Getenv("PORT")
	devicesHost := os.Getenv("DEVICES_HOST")
	usersHost := os.Getenv("USERS_HOST")
	env := os.Getenv("ENV")

	devicesClient, err := devices.NewDevicesClient(devicesHost)
	if err != nil {
		log.Fatal(err)
	}

	usersClient, err := users.NewUsersClient(usersHost)
	if err != nil {
		log.Fatal(err)
	}

	ginAuth, err := ginauth.NewGinAuth(usersClient)
	if err != nil {
		log.Fatal(err)
	}

	apiGateway := apigateway.NewAPIGateway(env, ginAuth, devicesClient, usersClient)
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
