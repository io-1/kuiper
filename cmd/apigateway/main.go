package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/io-1/kuiper/internal/apigateway"
	"github.com/io-1/kuiper/internal/logger/logruslogger"

	ginauth "github.com/io-1/kuiper/internal/apigateway/auth/ginauth"
	devices "github.com/io-1/kuiper/internal/apigateway/clients/devices"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions"
	users "github.com/io-1/kuiper/internal/apigateway/clients/users"
)

func main() {
	logger := logruslogger.NewLogrusLogger(true)

	port := os.Getenv("PORT")
	devicesHost := os.Getenv("DEVICES_HOST")
	usersHost := os.Getenv("USERS_HOST")
	interactionsHost := os.Getenv("INTERACTIONS_HOST")
	env := os.Getenv("ENV")

	devicesClient, err := devices.NewDevicesClient(devicesHost, logger)
	if err != nil {
		logger.Fatal(err)
	}

	usersClient, err := users.NewUsersClient(usersHost, logger)
	if err != nil {
		logger.Fatal(err)
	}

	interactionsClient, err := interactions.NewInteractionsClient(interactionsHost, logger)
	if err != nil {
		logger.Fatal(err)
	}

	ginAuth, err := ginauth.NewGinAuth(usersClient, logger)
	if err != nil {
		logger.Fatal(err)
	}

	apiGateway := apigateway.NewAPIGateway(env, ginAuth, devicesClient, usersClient, interactionsClient)
	router := gin.Default()

	err = apiGateway.InitV1Routes(router)
	if err != nil {
		logger.Fatal(err)
	}

	routerPort := fmt.Sprintf(":%s", port)
	logger.Infof("Listening on port: %s\n", port)
	err = apiGateway.Run(router, routerPort)
	if err != nil {
		logger.Fatal(err)
	}
}
