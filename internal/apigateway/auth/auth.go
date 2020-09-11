package auth

import "github.com/gin-gonic/gin"

type Auth interface {
	UseAuthMiddleware(c *gin.Context)
	LoginHandler(c *gin.Context)
	LogoutHandler(c *gin.Context)
	RefreshTokenHandler(c *gin.Context)
}
