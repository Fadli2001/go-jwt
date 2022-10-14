package middleware

import (
	"fmt"
	"go-jwt/utils/authenticator"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	acctToken authenticator.AccessToken
}

func (a *authTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
		fmt.Println(tokenString)
		if tokenString == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		
		token, err := a.acctToken.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		if token != nil {
			c.Next()
		}else {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
	}
}

func NewTokenValidator(acctToken authenticator.AccessToken) AuthTokenMiddleware {
	return &authTokenMiddleware{
		acctToken: acctToken,
	}
}
