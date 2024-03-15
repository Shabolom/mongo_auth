package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"test_hh_1/init/models"
	"test_hh_1/init/service"
	"test_hh_1/tools"
)

type UserAPI struct {
}

func NewUserApi() *UserAPI {
	return &UserAPI{}
}

var userService = service.NewUserService()

func (ua *UserAPI) GiveToken(c *gin.Context) {
	var user models.User

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err, accessToken, refToken := tools.CreateTokens()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err, code := userService.GiveToken(refToken, user)
	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.Writer.Header().Set("Access-Token", accessToken)
	c.Writer.Header().Set("Refresh-Token", refToken)

	c.JSON(code, "токены выданы")
}

func (ua *UserAPI) RefreshToken(c *gin.Context) {
	err, accessToken, refToken := tools.CreateTokens()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	result := userService.RefreshToken(c.Request.Header.Get("Refresh-Token"), refToken)

	c.Writer.Header().Set("Access-Token", accessToken)
	c.Writer.Header().Set("Refresh-Token", refToken)

	c.JSON(result.Status, result.Result)
}
