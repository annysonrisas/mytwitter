package jwt

import (
	"errors"
	"strings"

	"github.com/anny/mytwitter/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string, jwtsign string) (*models.Claim, bool, string, error) {
	miClave := []byte(jwtsign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Formato de token invalido")

	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		//Rutina que chequea con la BD
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Invalido")
	}
	return &claims, false, string(""), err

}
