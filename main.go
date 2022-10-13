package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}


type Credential struct {
	Username string
	Password string
}


func main(){
	router := gin.Default()
	router.Use(AuthTokenMiddleware())
	router.POST("/login", func(c *gin.Context) {
		var user Credential
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "can't bind struct",
			})
			return
		}
		if user.Username == "enigma" && user.Password == "123" {
			c.JSON(200,gin.H{
				"token" : "123",
			})
		}else {
			c.AbortWithStatus(401)
		}
	})

	router.GET("/customer",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message" : "customer",
		})
	})

	err := router.Run(":8888")
	if err != nil {
		panic(err)
	}
}

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			c.Next()
		}else{
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(401, gin.H {
					"message" : "Unauthorized",
				})
				c.Abort()
			}
			if h.AuthorizationHeader == "123" {
				c.Next()
			}else{
				c.JSON(401, gin.H {
					"message" : "Unauthorized",
				})
				c.Abort()
			}
		}
	}
}