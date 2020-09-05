package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt"
	devices "github.com/n7down/kuiper/internal/apigateway/clients/devices"
	users "github.com/n7down/kuiper/internal/apigateway/clients/users"
)

type APIGateway struct {
	authMiddleware *jwt.GinJWTMiddleware
	devicesClient  *devices.DevicesClient
	usersClient    *users.UsersClient
}

func NewAPIGateway(ginJWTMiddleware *jwt.GinJWTMiddleware, devicesClient *devices.DevicesClient, usersClient *users.UsersClient) *APIGateway {
	return &APIGateway{
		authMiddleware: ginJWTMiddleware,
		devicesClient:  devicesClient,
		usersClient:    usersClient,
	}
}

func (g *APIGateway) InitV1Routes(r *gin.Engine) error {
	v1 := r.Group("/api/v1")

	// FIXME: take out middlewrae for error
	v1.POST("/login", g.authMiddleware.LoginHandler)
	v1.GET("/refresh_token", g.authMiddleware.RefreshHandler)

	authGroup := v1.Group("/auth")
	authGroup.Use(g.authMiddleware.MiddlewareFunc())
	{
		authGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				// FIXME: get user info from the claims
				"text": "Hello World.",
			})
		})
	}

	deviceGroup := v1.Group("/devices")
	{
		deviceGroup.POST("/bc", g.devicesClient.CreateBatCaveDeviceSetting)
		deviceGroup.GET("/bc/:device_id", g.devicesClient.GetBatCaveDeviceSetting)
		deviceGroup.PUT("/bc/:device_id", g.devicesClient.UpdateBatCaveDeviceSetting)
	}

	usersGroup := v1.Group("/users")
	{
		usersGroup.POST("/create", g.usersClient.CreateUser)
		usersGroup.GET("/:username", g.usersClient.GetUser)
		usersGroup.PUT("/:username", g.usersClient.UpdateUser)
		usersGroup.DELETE("/:username", g.usersClient.DeleteUser)
	}

	r.NoRoute(g.authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return nil
}

func (g *APIGateway) Run(router *gin.Engine, port string) error {
	err := http.ListenAndServe(port, router)
	if err != nil {
		return err
	}
	return nil
}
