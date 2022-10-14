package controller

import (
	"go-jwt/delivery/middleware"
	"go-jwt/model"
	"go-jwt/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AppController struct {
	rg *gin.RouterGroup
	authUseCase usecase.AuthUseCase
}

func (cc *AppController) userAuth(ctx *gin.Context) {
	var user model.UserCredential
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"can't bind struct",
		})
		return
	}

	token, err := cc.authUseCase.UserAuth(user)
	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"token":token,
	})
}

func (cc *AppController) getCustomer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,gin.H{
		"message" : "user",
	})
}

func NewAppController(routerGroup *gin.RouterGroup, authUseCase usecase.AuthUseCase,tokenMdw middleware.AuthTokenMiddleware) *AppController {
	controller := AppController {
		rg: routerGroup,
		authUseCase: authUseCase,
	}
	
	controller.rg.POST("/auth",controller.userAuth)
	protectedGroup := controller.rg.Group("/protected",tokenMdw.RequireToken())
	protectedGroup.GET("/user",controller.getCustomer)
	return &controller
}