package main

import (
	"github.com/gin-gonic/gin"
)

type authHeader struct{
	AuthorizationHeader string `header:"Authorization"`
}

func main() {
	router := gin.Default()

	router.GET("/customer",func(ctx *gin.Context) {
		h := authHeader{}
		err := ctx.ShouldBindHeader(&h)
		if err != nil {
			ctx.JSON(401,gin.H{
				"message":"Unauthorized",
			})
			return
		}
		if h.AuthorizationHeader == "123" {
			ctx.JSON(200,gin.H{
				"message":"customer",
			})	
			return		
		}
		ctx.JSON(401,gin.H{
			"message" : "Unauthorized",
		}) 
	})

	err := router.Run(":8888")
	if err != nil {
		panic(err)
	}	

}
