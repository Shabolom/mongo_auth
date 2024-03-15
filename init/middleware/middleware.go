package middleware

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test_hh_1/tools"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		checkToken := base64.StdEncoding.EncodeToString([]byte(c.Request.Header.Get("Access-Token")))

		fmt.Println(checkToken == c.Request.Header.Get("Access-Token"))

		if checkToken != c.Request.Header.Get("Refresh-Token") {
			tools.CreateError(http.StatusBadRequest, errors.New("токены не соответствуют"), c)
			return
		}

		c.Next()
	}
}
