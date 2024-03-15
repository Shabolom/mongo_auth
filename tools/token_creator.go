package tools

import (
	"encoding/base64"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"test_hh_1/config"
	"time"
)

type Climes struct {
	jwt.RegisteredClaims
	Token  string
	UserID uuid.UUID
}

func AccessToken() BaseResult {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		Issuer:    "Timur",
		Subject:   "user",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	strToken, err := token.SignedString([]byte(config.Env.SecretKey))
	if err != nil {
		return BaseResult{
			Err:    err,
			Status: http.StatusInternalServerError,
			Result: "",
		}
	}

	return BaseResult{
		Result: strToken,
	}
}

// CreateTokens  возвращает error, Access-Token и Refresh-Token
func CreateTokens() (error, string, string) {
	accessToken := AccessToken()

	if accessToken.Err != nil {
		return accessToken.Err, "", ""
	}

	refToken := base64.StdEncoding.EncodeToString([]byte(accessToken.Result))

	return nil, accessToken.Result, refToken
}
