package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	devices "github.com/n7down/kuiper/internal/apigateway/clients/devices"
	"github.com/n7down/kuiper/internal/apigateway/clients/users"
)

type APIGateway struct {
	devicesClient *devices.DevicesClient
	usersClient   *users.UsersClient
}

func NewAPIGateway(c *devices.DevicesClient, u *users.UsersClient) *APIGateway {
	return &APIGateway{
		devicesClient: c,
		usersClient:   u,
	}
}

func (g *APIGateway) InitV1Routes(r *gin.Engine) error {
	v1 := r.Group("/api/v1")
	deviceGroup := v1.Group("/devices")
	{
		deviceGroup.POST("/bc", g.devicesClient.CreateBatCaveDeviceSetting)
		deviceGroup.GET("/bc/:device_id", g.devicesClient.GetBatCaveDeviceSetting)
		deviceGroup.PUT("/bc/:device_id", g.devicesClient.UpdateBatCaveDeviceSetting)
	}

	usersGroup := v1.Group("/users")
	{
		usersGroup.POST("/create", nil)
		usersGroup.GET("/:user_id", nil)
		usersGroup.PUT("/:user_id", nil)
		usersGroup.DELETE("/:user_id", nil)
	}

	return nil
}

func (g *APIGateway) Run(router *gin.Engine, port string) error {
	err := http.ListenAndServe(port, router)
	if err != nil {
		return err
	}
	return nil
}
