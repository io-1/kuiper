package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/n7down/kuiper/internal/apigateway"
	"github.com/n7down/kuiper/internal/logger/logruslogger"

	"github.com/n7down/kuiper/internal/apigateway/auth/ginauth"
	devices "github.com/n7down/kuiper/internal/apigateway/clients/devices"
	"github.com/n7down/kuiper/internal/apigateway/clients/users"
	log "github.com/sirupsen/logrus"
)

func init() {
}

func main() {
	log.SetReportCaller(true)

	port := os.Getenv("PORT")
	devicesHost := os.Getenv("DEVICES_HOST")
	usersHost := os.Getenv("USERS_HOST")

	logger := logruslogger.NewLogrusLogger(true)

	devicesClient, err := devices.NewDevicesClient(devicesHost, logger)
	if err != nil {
		log.Fatal(err)
	}

	usersClient, err := users.NewUsersClient(usersHost)
	if err != nil {
		log.Fatal(err)
	}

	ginAuth := ginauth.NewGinAuth(usersClient)
	authMiddleware, err := ginAuth.GetAuthMiddleware()
	if err != nil {
		log.Fatal(err)
	}

	apiGateway := apigateway.NewAPIGateway(authMiddleware, devicesClient, usersClient)
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
