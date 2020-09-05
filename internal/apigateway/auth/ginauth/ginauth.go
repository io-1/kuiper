package ginauth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/kuiper/internal/apigateway/auth/request"
	"github.com/n7down/kuiper/internal/apigateway/auth/response"
	"github.com/n7down/kuiper/internal/apigateway/clients/users"
	"github.com/n7down/kuiper/internal/utils"

	jwt "github.com/appleboy/gin-jwt"
)

const (
	identityKey = "id"
)

var (
	loginResponse *response.LoginResponse
)

type GinAuth struct {
	usersClient *users.UsersClient
}

func NewGinAuth(u *users.UsersClient) *GinAuth {
	return &GinAuth{
		usersClient: u,
	}
}

func (a *GinAuth) GetAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "dev",
		Key:         []byte("43deio1"),
		Timeout:     time.Duration(3) * time.Hour,
		MaxRefresh:  time.Duration(3) * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*response.LoginResponse); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					"username":  v.Username,
					"name":      v.Name,
					"email":     v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &response.LoginResponse{
				ID: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var (
				req request.LoginRequest
			)

			if err := c.ShouldBind(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if validationErrors := req.Validate(); len(validationErrors) > 0 {
				return "", jwt.ErrFailedAuthentication
			}

			// get the user
			res, err := a.usersClient.GetUserLogin(req.Username)
			if err != nil {
				return nil, err
			}

			if res.Username == "" {
				return nil, jwt.ErrMissingLoginValues
			}

			// check if password is valid with bcrypt
			isValidPassword, err := utils.CheckUserSecret(res.Password, req.Password)
			if err != nil {
				return nil, err
			}

			if isValidPassword {
				loginResponse = &response.LoginResponse{
					ID:       res.ID,
					Username: res.Username,
					Name:     res.Name,
					Email:    res.Email,
				}
				return loginResponse, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// TODO: use this to see if the user has access to the route - used to implement permissions
			// if v, ok := data.(*response.LoginResponse); ok && v.ID == "some-id" {
			// 	return true
			// }

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, statusCode int, token string, tokenExpires time.Time) {
			loginResponse.Token = token
			loginResponse.Expires = tokenExpires
			c.JSON(http.StatusOK, loginResponse)
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
