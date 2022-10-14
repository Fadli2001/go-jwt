package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	ApiConfig
	TokenConfig
}

func (c Config) readConfig() Config {
	c.ApiConfig = ApiConfig{
		ApiPort: "8889",
		ApiHost: "localhost",
	}
	c.TokenConfig = TokenConfig{
		ApplicationName:     "ENIGMA",
		JwtSignatureKey:     "P@ssword",
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: 60 * time.Second,
	}
	return c
}

func NewConfig() Config {
	conf := Config{}
	return conf.readConfig()
	
}