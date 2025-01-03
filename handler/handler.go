package handler

import (
	"context"
	"fmt"

	"github.com/anny/mytwitter/jwt"
	"github.com/anny/mytwitter/models"
	"github.com/anny/mytwitter/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {

	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + "> " + ctx.Value(models.Key("method")).(string))
	var r models.RespApi
	r.Status = 400

	isOk, statusCode, msg, _ := validoAuthorizacion(ctx, request)
	if !isOk {
		r.Status = statusCode
		r.Message = msg
		return r
	}
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)

		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//

	}
	r.Message = "Method Invalid"
	return r

}

func validoAuthorizacion(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}
	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}
	claim, todoOk, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOk {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, err.Error(), models.Claim{}
		}
	}
	fmt.Println("Token Ok")
	return true, 200, msg, *claim
}
