package routes

import (
	"github.com/gin-gonic/gin"
	"test_hh_1/init/api"
	"test_hh_1/init/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	user := api.NewUserApi()

	authRequired := r.Group("/")
	authRequired.Use(middleware.CheckToken())

	r.GET("api/get_tokens", user.GiveToken)

	authRequired.GET("api/refresh_tokens", user.RefreshToken)

	return r
}
