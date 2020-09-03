package apigateway

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/kuiper/internal/apigateway/clients/users"

	jwt "github.com/appleboy/gin-jwt"
	devices "github.com/n7down/kuiper/internal/apigateway/clients/devices"
)

// type login struct {
// 	Username string `form:"username" json:"username" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

// var identityKey = "id"

// func helloHandler(c *gin.Context) {
// 	claims := jwt.ExtractClaims(c)
// 	user, _ := c.Get(identityKey)
// 	c.JSON(200, gin.H{
// 		"userID":   claims[identityKey],
// 		"userName": user.(*User).Username,
// 		"text":     "Hello World.",
// 	})
// }

// User demo
type User struct {
	// UserName  string
	// FirstName string
	// LastName  string
	Username string
	Name     string
	Email    string
}

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResponse struct {
	Username string
	Name     string
	Email    string
	Token    string
	Expires  string
}

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

func (g *APIGateway) authRoutes() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "dev",
		Key:         []byte("43deio1"),
		Timeout:     time.Duration(3) * time.Hour,
		MaxRefresh:  time.Duration(3) * time.Hour,
		IdentityKey: "username",
		// PayloadFunc: func(data interface{}) jwt.MapClaims {
		// 	if v, ok := data.(*User); ok {
		// 		return jwt.MapClaims{
		// 			identityKey: v.UserName,
		// 		}
		// 	}
		// 	return jwt.MapClaims{}
		// },
		// IdentityHandler: func(c *gin.Context) interface{} {
		// 	claims := jwt.ExtractClaims(c)
		// 	return &User{
		// 		UserName: claims[identityKey].(string),
		// 	}
		// },
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login LoginRequest
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			// userID := login.Username
			// password := login.Password

			if login.Username == "admin" && login.Password == "admin" {
				return &User{
					// 	UserName:  userID,
					// 	LastName:  "Bo-Yi",
					// 	FirstName: "Wu",
					Username: "admin",
					Name:     "admin",
					Email:    "admin@io1.com",
				}, nil
			}

			// FIXME: get the user
			// FIXME: check the password with the util bcrypt - CheckUserSecret

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// if v, ok := data.(*User); ok && v.UserName == "admin" {
			// 	return true
			// }
			// return false
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, statusCode int, token string, tokenExpires time.Time) {
			// FIXME: add this to the response
			// authLoginResponse.Token = token
			// authLoginResponse.Expires = tokenExpires
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		return authMiddleware, err
	}

	err = authMiddleware.MiddlewareInit()
	if err != nil {
		return authMiddleware, err
	}

	return authMiddleware, nil
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
		usersGroup.POST("/create", g.usersClient.CreateUser)
		usersGroup.GET("/:username", g.usersClient.GetUser)
		usersGroup.PUT("/:username", g.usersClient.UpdateUser)
		usersGroup.DELETE("/:username", g.usersClient.DeleteUser)
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
