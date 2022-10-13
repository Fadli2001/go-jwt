package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/customer",func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{
			"message" : "customer",
		})
	})
	err := router.Run(":8888")
	if err != nil {
		panic(err)
	}
}