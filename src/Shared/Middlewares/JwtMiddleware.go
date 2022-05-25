package Middlewares

import (
	"TemplateSolution/src/Shared/Utils"
	"github.com/labstack/echo/v4/middleware"
)

var signingKey = []byte(Utils.SETTING.SecreyKey)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: signingKey,
})
