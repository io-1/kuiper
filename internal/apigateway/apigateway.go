package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"

	"github.com/io-1/kuiper/internal/apigateway/auth/ginauth"

	"github.com/io-1/kuiper/internal/apigateway/clients/devicesclient"
	"github.com/io-1/kuiper/internal/apigateway/clients/interactionsclient"
	"github.com/io-1/kuiper/internal/apigateway/clients/usersclient"

	jwt "github.com/appleboy/gin-jwt"
)

const (
	DEV = "dev"
)

type APIGateway struct {
	env                string
	version            string
	build              string
	ginAuth            *ginauth.GinAuth
	devicesClient      *devicesclient.DevicesClient
	usersClient        *usersclient.UsersClient
	interactionsClient *interactionsclient.InteractionsClient
}

func NewAPIGateway(env, version, build string, ginAuth *ginauth.GinAuth, devicesClient *devicesclient.DevicesClient, usersClient *usersclient.UsersClient, interactionsClient *interactionsclient.InteractionsClient) *APIGateway {
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

	devicesGroup := v1.Group("/devices")
	{

		// settings are stored in a database
		settingsGroup := devicesGroup.Group("/settings")
		{

			humidityGroup := settingsGroup.Group("/humidity")
			{
				// FIXME: change to moister devices or something similiar
				humidityGroup.POST("", g.devicesClient.CreateBatCaveDeviceSetting)
				humidityGroup.GET("/:id", g.devicesClient.GetBatCaveDeviceSetting)
				humidityGroup.PUT("/:id", g.devicesClient.UpdateBatCaveDeviceSetting)
			}
		}

		// send - just sends out a command
		sendGroup := devicesGroup.Group("/send")
		{
			lampGroup := sendGroup.Group("/lamp")
			{
				lampGroup.POST("/:send_lamp_mac/on", g.devicesClient.SendLampDeviceOn)
				lampGroup.POST("/:send_lamp_mac/off", g.devicesClient.SendLampDeviceOff)
				lampGroup.POST("/:send_lamp_mac/toggle", g.devicesClient.SendLampDeviceToggle)
				lampGroup.POST("/:send_lamp_mac/color", g.devicesClient.SendLampDeviceColor)
				lampGroup.POST("/:send_lamp_mac/brightness", g.devicesClient.SendLampDeviceBrightness)
				lampGroup.POST("/:send_lamp_mac/brightness/auto/on", g.devicesClient.SendLampDeviceAutoBrightnessOn)
				lampGroup.POST("/:send_lamp_mac/brightness/auto/off", g.devicesClient.SendLampDeviceAutoBrightnessOff)
				lampGroup.POST("/:send_lamp_mac/brightness/auto/toggle", g.devicesClient.SendLampDeviceAutoBrightnessToggle)
				lampGroup.POST("/:send_lamp_mac/pulse", g.devicesClient.SendLampDevicePulse)
			}
		}
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
		interactionsGroup.GET("/:interaction_id/details", g.interactionsClient.GetInteractionDetails)

	}

	keypadGroup := v1.Group("/keypad")
	{
		conditionsGroup := keypadGroup.Group("/condition")
		{
			conditionsGroup.POST("", g.interactionsClient.CreateKeypadCondition)
			conditionsGroup.GET("/:keypad_condition_id", g.interactionsClient.GetKeypadCondition)
			conditionsGroup.PUT("/:keypad_condition_id", g.interactionsClient.UpdateKeypadCondition)
			conditionsGroup.PATCH("/:keypad_condition_id", g.interactionsClient.PatchKeypadCondition)
			conditionsGroup.DELETE("/:keypad_condition_id", g.interactionsClient.DeleteKeypadCondition)
		}
	}

	lampGroup := v1.Group("/lamp")
	{
		eventGroup := lampGroup.Group("/event")
		{
			onGroup := eventGroup.Group("/on")
			{
				onGroup.POST("", g.interactionsClient.CreateLampOnEvent)
				onGroup.GET("/:lamp_on_event_id", g.interactionsClient.GetLampOnEvent)
				onGroup.PUT("/:lamp_on_event_id", g.interactionsClient.UpdateLampOnEvent)
				onGroup.PATCH("/:lamp_on_event_id", g.interactionsClient.PatchLampOnEvent)
				onGroup.DELETE("/:lamp_on_event_id", g.interactionsClient.DeleteLampOnEvent)
			}

			offGroup := eventGroup.Group("/off")
			{
				offGroup.POST("", g.interactionsClient.CreateLampOffEvent)
				offGroup.GET("/:lamp_off_event_id", g.interactionsClient.GetLampOffEvent)
				offGroup.PUT("/:lamp_off_event_id", g.interactionsClient.UpdateLampOffEvent)
				offGroup.PATCH("/:lamp_off_event_id", g.interactionsClient.PatchLampOffEvent)
				offGroup.DELETE("/:lamp_off_event_id", g.interactionsClient.DeleteLampOffEvent)
			}

			toggleGroup := eventGroup.Group("/toggle")
			{
				toggleGroup.POST("", g.interactionsClient.CreateLampToggleEvent)
				toggleGroup.GET("/:lamp_toggle_event_id", g.interactionsClient.GetLampToggleEvent)
				toggleGroup.PUT("/:lamp_toggle_event_id", g.interactionsClient.UpdateLampToggleEvent)
				toggleGroup.PATCH("/:lamp_toggle_event_id", g.interactionsClient.PatchLampToggleEvent)
				toggleGroup.DELETE("/:lamp_toggle_event_id", g.interactionsClient.DeleteLampToggleEvent)
			}

			// FIXME: finished adding grpc calls - need to add this section
			brightnessGroup := eventGroup.Group("/brightness")
			{
				brightnessGroup.POST("", g.interactionsClient.CreateLampBrightnessEvent)
				brightnessGroup.GET("/:lamp_brightness_event_id", g.interactionsClient.GetLampBrightnessEvent)
				brightnessGroup.PUT("/:lamp_brightness_event_id", g.interactionsClient.UpdateLampBrightnessEvent)
				brightnessGroup.PATCH("/:lamp_brightness_event_id", g.interactionsClient.PatchLampBrightnessEvent)
				brightnessGroup.DELETE("/:lamp_brightness_event_id", g.interactionsClient.DeleteLampBrightnessEvent)

				autoBrightnessGroup := eventGroup.Group("/auto")
				{
					onGroup := autoBrightnessGroup.Group("/on")
					{
						onGroup.POST("", g.interactionsClient.CreateLampAutoBrightnessOnEvent)
						onGroup.GET("/:lamp_auto_brightness_on_event_id", g.interactionsClient.GetLampAutoBrightnessOnEvent)
						onGroup.PUT("/:lamp_auto_brightness_on_event_id", g.interactionsClient.UpdateLampAutoBrightnessOnEvent)
						onGroup.PATCH("/:lamp_auto_brightness_on_event_id", g.interactionsClient.PatchLampAutoBrightnessOnEvent)
						onGroup.DELETE("/:lamp_auto_brightness_on_event_id", g.interactionsClient.DeleteLampAutoBrightnessOnEvent)
					}

					offGroup := autoBrightnessGroup.Group("/off")
					{
						offGroup.POST("", g.interactionsClient.CreateLampAutoBrightnessOffEvent)
						offGroup.GET("/:lamp_auto_brightness_off_event_id", g.interactionsClient.GetLampAutoBrightnessOffEvent)
						offGroup.PUT("/:lamp_auto_brightness_off_event_id", g.interactionsClient.UpdateLampAutoBrightnessOffEvent)
						offGroup.PATCH("/:lamp_auto_brightness_off_event_id", g.interactionsClient.PatchLampAutoBrightnessOffEvent)
						offGroup.DELETE("/:lamp_auto_brightness_off_event_id", g.interactionsClient.DeleteLampAutoBrightnessOffEvent)
					}

					toggleGroup := autoBrightnessGroup.Group("/toggle")
					{
						toggleGroup.POST("", g.interactionsClient.CreateLampAutoBrightnessToggleEvent)
						toggleGroup.GET("/:lamp_auto_brightness_toggle_event_id", g.interactionsClient.GetLampAutoBrightnessToggleEvent)
						toggleGroup.PUT("/:lamp_auto_brightness_toggle_event_id", g.interactionsClient.UpdateLampAutoBrightnessToggleEvent)
						toggleGroup.PATCH("/:lamp_auto_brightness_toggle_event_id", g.interactionsClient.PatchLampAutoBrightnessToggleEvent)
						toggleGroup.DELETE("/:lamp_auto_brightness_toggle_event_id", g.interactionsClient.DeleteLampAutoBrightnessToggleEvent)
					}
				}
			}

			colorGroup := eventGroup.Group("/color")
			{
				colorGroup.POST("", g.interactionsClient.CreateLampColorEvent)
				colorGroup.GET("/:lamp_color_event_id", g.interactionsClient.GetLampColorEvent)
				colorGroup.PUT("/:lamp_color_event_id", g.interactionsClient.UpdateLampColorEvent)
				colorGroup.PATCH("/:lamp_color_event_id", g.interactionsClient.PatchLampColorEvent)
				colorGroup.DELETE("/:lamp_color_event_id", g.interactionsClient.DeleteLampColorEvent)
			}

			pulseGroup := eventGroup.Group("/pulse")
			{
				pulseGroup.POST("", g.interactionsClient.CreateLampPulseEvent)
				pulseGroup.GET("/:lamp_pulse_event_id", g.interactionsClient.GetLampPulseEvent)
				pulseGroup.PUT("/:lamp_pulse_event_id", g.interactionsClient.UpdateLampPulseEvent)
				pulseGroup.PATCH("/:lamp_pulse_event_id", g.interactionsClient.PatchLampPulseEvent)
				pulseGroup.DELETE("/:lamp_pulse_event_id", g.interactionsClient.DeleteLampPulseEvent)
			}
		}
	}

	interactGroup := v1.Group("/interact")
	{

		keypadConditionToLampEventGroup := interactGroup.Group("/keypad/lamp")
		{
			keypadConditionToLampEventGroup.POST("", g.interactionsClient.CreateKeypadConditionToLampEventInteraction)
			keypadConditionToLampEventGroup.GET("/:keypad_to_lamp_id", g.interactionsClient.GetKeypadConditionToLampEventInteraction)
			keypadConditionToLampEventGroup.PUT("/:keypad_to_lamp_id", g.interactionsClient.UpdateKeypadConditionToLampEventInteraction)
			keypadConditionToLampEventGroup.PATCH("/:keypad_to_lamp_id", g.interactionsClient.PatchKeypadConditionToLampEventInteraction)
			keypadConditionToLampEventGroup.DELETE("/:keypad_to_lamp_id", g.interactionsClient.DeleteKeypadConditionToLampEventInteraction)
		}
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
