package handler

import (
	"errors"
	"omo-msa-account/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

func tokenFromJWT(_userid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(config.Schema.Token.JWT.Expiry)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = _userid
	token.Claims = claims
	return token.SignedString([]byte(config.Schema.Encrypt.Secret))
}

func useridFromToken(_token string, _strategy proto.Strategy) (string, error) {
	if proto.Strategy_STRATEGY_JWT == _strategy {
		token, err := jwt.Parse(_token, func(_t *jwt.Token) (interface{}, error) {
			return []byte(config.Schema.Encrypt.Secret), nil
		})
		if nil != err {
			return "", err
		}

		if !token.Valid {
			return "", errors.New("Token is not valid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return "", errors.New("failure to convert Claims")
		}
		return claims["id"].(string), nil
	}

	return _token, nil
}
