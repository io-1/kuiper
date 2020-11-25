package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"

	"github.com/io-1/kuiper/internal/apigateway/auth/ginauth"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactions"

	jwt "github.com/appleboy/gin-jwt"
	devices "github.com/io-1/kuiper/internal/apigateway/clients/devices"
	users "github.com/io-1/kuiper/internal/apigateway/clients/users"
)

const (
	DEV = "dev"
)

type APIGateway struct {
	env                string
	version            string
	build              string
	ginAuth            *ginauth.GinAuth
	devicesClient      *devices.DevicesClient
	usersClient        *users.UsersClient
	interactionsClient *interactions.InteractionsClient
}

func NewAPIGateway(env, version, build string, ginAuth *ginauth.GinAuth, devicesClient *devices.DevicesClient, usersClient *users.UsersClient, interactionsClient *interactions.InteractionsClient) *APIGateway {
	return &APIGateway{
		env:                env,
		version:            version,
		build:              build,
		ginAuth:            ginAuth,
		devicesClient:      devicesClient,
		usersClient:        usersClient,
		interactionsClient: interactionsClient,
	}
}

func (g *APIGateway) wrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (g *APIGateway) InitV1Routes(r *gin.Engine) error {
	if g.env == DEV {
		opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
		sh := middleware.Redoc(opts, nil)
		r.GET("/swagger.yaml", func(c *gin.Context) {
			c.File("api/swagger.yaml")
		})
		r.GET("/docs", g.wrapH(sh))

		r.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version": g.version,
				"build":   g.build,
			})
		})
	}

	v1 := r.Group("/api/v1")
	v1.POST("/login", g.ginAuth.LoginHandler)

	authGroup := v1.Group("/auth")
	authGroup.Use(g.ginAuth.UseAuthMiddleware)
	{
		authGroup.GET("/refresh_token", g.ginAuth.RefreshTokenHandler)

		authGroup.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			c.JSON(http.StatusOK, gin.H{
				"id":       claims["id"],
				"username": claims["username"],
				"name":     claims["name"],
				"email":    claims["email"],
				"text":     "Hello World.",
			})
		})

		authGroup.POST("/logout", g.ginAuth.LogoutHandler)
	}

	deviceGroup := v1.Group("/devices")
	{
		deviceGroup.POST("/bc", g.devicesClient.CreateBatCaveDeviceSetting)
		deviceGroup.GET("/bc/:id", g.devicesClient.GetBatCaveDeviceSetting)
		deviceGroup.PUT("/bc/:id", g.devicesClient.UpdateBatCaveDeviceSetting)
	}

	usersGroup := v1.Group("/users")
	{
		usersGroup.POST("", g.usersClient.CreateUser)
		usersGroup.GET("/:user_id", g.usersClient.GetUser)
		usersGroup.PUT("/:user_id", g.usersClient.UpdateUser)
		usersGroup.PATCH("/:user_id", g.usersClient.PatchUser)
		usersGroup.DELETE("/:user_id", g.usersClient.DeleteUser)
	}

	interactionsGroup := v1.Group("/interactions")
	{
		interactionsGroup.POST("", g.interactionsClient.CreateInteraction)
		interactionsGroup.GET("/:interaction_id", g.interactionsClient.GetInteraction)
		interactionsGroup.PUT("/:interaction_id", g.interactionsClient.UpdateInteraction)
		interactionsGroup.PATCH("/:interaction_id", g.interactionsClient.PatchInteraction)
		interactionsGroup.DELETE("/:interaction_id", g.interactionsClient.DeleteInteraction)

	}

	conditionsGroup := v1.Group("/conditions")
	{
		keypadConditionsGroup := conditionsGroup.Group("/keypad")
		{
			keypadConditionsGroup.POST("", g.interactionsClient.CreateKeypadCondition)
			keypadConditionsGroup.GET("/:keypad_id", g.interactionsClient.GetKeypadCondition)
			keypadConditionsGroup.PUT("/:keypad_id", g.interactionsClient.UpdateKeypadCondition)
			keypadConditionsGroup.PATCH("/:keypad_id", g.interactionsClient.PatchKeypadCondition)
			keypadConditionsGroup.DELETE("/:keypad_id", g.interactionsClient.DeleteKeypadCondition)
		}
	}

	eventsGroup := v1.Group("/events")
	{
		lampEventGroup := eventsGroup.Group("/lamp")
		{
			lampEventGroup.POST("", g.interactionsClient.CreateLampEvent)
			lampEventGroup.GET("/:lamp_id", g.interactionsClient.GetLampEvent)
			lampEventGroup.PUT("/:lamp_id", g.interactionsClient.UpdateLampEvent)
			lampEventGroup.PATCH("/:lamp_id", g.interactionsClient.PatchLampEvent)
			lampEventGroup.DELETE("/:lamp_id", g.interactionsClient.DeleteLampEvent)
		}
	}

	// attach conditions to events
	attachGroup := v1.Group("/attach")
	{
		attachGroup.POST("", g.interactionsClient.CreateAttach)
		attachGroup.GET("/:attach_id", g.interactionsClient.GetAttach)
		attachGroup.PUT("/:attach_id", g.interactionsClient.UpdateAttach)
		attachGroup.PATCH("/:attach_id", g.interactionsClient.PatchAttach)
		attachGroup.DELETE("/:attach_id", g.interactionsClient.DeleteAttach)
	}

	r.NoRoute(g.ginAuth.UseAuthMiddleware, func(c *gin.Context) {
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
