package handlers

import (
	"context"
	"fmt"
	"go/token"

	"github.com/aws/aws-lambda-go/events"
	"github.com/luiscomas/liketwitterWithGo/jwt"
	"github.com/luiscomas/liketwitterWithGo/models"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespAPI {
	fmt.Println("procesando " + ctx.Value(models.Key("path")).(string) + " " + ctx.Value(models.Key("method")).(string))

	var r models.RespAPI
	r.Status = 400

	isOK, statusCode, message, clain := ValidateToken(ctx, request)

	switch ctx.Value(models.Key("method")).(string) {
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		}
	}

	r.Message = "Method not Invalid"
	return r
}
 func ValidateToken(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "login" || path == "signup" || path == "getAvatar" || path == "getBanner"{
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 400, "Token requerido", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcesarToken(token, ctx.Value(models.Key("JWTSign")).(string))
}