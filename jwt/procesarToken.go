package jwt

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/luiscomas/liketwitterWithGo/models"
)

var Email string
var IDUser string


func ProcesarToken (token string, JWTSign string) (*models.Claim, bool,string,error) {
    mykey := []byte(JWTSign)
    var claims models.Claim
    splitToken := strings.Split(token, "Bearer")
    if len(splitToken) != 2 {
        return &claims, false, string(""), errors.New("No valid token format")
    }
    token = strings.TrimSpace(splitToken[1])
    tkn , err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
        return mykey, nil
    })
    if err == nil {
        // Rutine check againth databse
    }

    if !tkn.Valid {
        return &claims, false, string(""),errors.New("Invalid token")

    }
    return &claims, false, string(""), err
}
