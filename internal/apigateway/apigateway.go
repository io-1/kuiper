package apigateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"

	jwt "github.com/appleboy/gin-jwt"
	devices "github.com/io-1/kuiper/internal/apigateway/clients/devices"
	users "github.com/io-1/kuiper/internal/apigateway/clients/users"
)

const (
	DEV = "dev"
)

type APIGateway struct {
	env            string
	authMiddleware *jwt.GinJWTMiddleware
	devicesClient  *devices.DevicesClient
	usersClient    *users.UsersClient
}

func NewAPIGateway(env string, ginJWTMiddleware *jwt.GinJWTMiddleware, devicesClient *devices.DevicesClient, usersClient *users.UsersClient) *APIGateway {
	return &APIGateway{
		env:            env,
		authMiddleware: ginJWTMiddleware,
		devicesClient:  devicesClient,
		usersClient:    usersClient,
	}
}

func (g *APIGateway) wrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// swagger:route POST /api/v1/login auth
//
// Login as a user.
//
// Allows a user to login.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: LoginResponse
func (g *APIGateway) loginHandler(c *gin.Context) {
	g.authMiddleware.LoginHandler(c)
}

// swagger:route POST /api/v1/auth/logout auth
//
// Logout as a user.
//
// Allows a user to logout.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: LogoutResponse
func (g *APIGateway) logoutHandler(c *gin.Context) {
	g.authMiddleware.LogoutHandler(c)
}

func (g *APIGateway) refreshTokenHandler(c *gin.Context) {
	g.authMiddleware.RefreshHandler(c)
}

func (g *APIGateway) InitV1Routes(r *gin.Engine) error {
	if g.env == DEV {
		opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
		sh := middleware.Redoc(opts, nil)
		r.GET("/swagger.yaml", func(c *gin.Context) {
			c.File("api/swagger.yaml")
		})
		r.GET("/docs", g.wrapH(sh))
	}

	v1 := r.Group("/api/v1")
	v1.POST("/login", g.loginHandler)

	authGroup := v1.Group("/auth")
	authGroup.Use(g.authMiddleware.MiddlewareFunc())
	{
		authGroup.GET("/refresh_token", g.refreshTokenHandler)

		authGroup.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			c.JSON(200, gin.H{
				"id":       claims["id"],
				"username": claims["username"],
				"name":     claims["name"],
				"email":    claims["email"],
				"text":     "Hello World.",
			})
		})

		authGroup.POST("/logout", g.logoutHandler)
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
