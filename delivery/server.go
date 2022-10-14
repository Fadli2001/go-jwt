package delivery

import (
	"fmt"
	"go-jwt/config"
	"go-jwt/delivery/controller"
	"go-jwt/delivery/middleware"
	"go-jwt/usecase"
	"go-jwt/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host 	string
	authUseCase usecase.AuthUseCase
	tokenService authenticator.AccessToken
}

func (s *Server) initController() {
	publicRoute := s.engine.Group("/enigma")
	tokenMdw := middleware.NewTokenValidator(s.tokenService)
	controller.NewAppController(publicRoute,s.authUseCase,tokenMdw)
}

func (s *Server) Run(){
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	c := config.NewConfig()
	r := gin.Default()
	tokenService := authenticator.NewAccessToken(c.TokenConfig)
	authUseCase := usecase.NewAuthUseCase(tokenService)	
	if c.ApiHost == "" || c.ApiPort == "" {
		panic("No Host or port Define")
	}
	host := fmt.Sprintf("%s:%s",c.ApiHost,c.ApiPort)

	return &Server{engine: r,host: host,authUseCase: authUseCase,tokenService: tokenService}
}