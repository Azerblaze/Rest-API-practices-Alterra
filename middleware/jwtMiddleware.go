package middleware

import (
	"praktikum/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtCustomClaims struct {
	Authorized bool `json:"authorized"`
	ID         int  `json:"id"`
	jwt.StandardClaims
}

func CreateToken(userId int) (string, error) {
	// claims := jwt.MapClaims{}

	// Set custom claims
	claims := &JwtCustomClaims{
		true,
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	// claims["authorized"] = true
	// claims["userId"] = userId
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires in one hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}

// func JWTMiddleware(e *echo.Echo) {
// 	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
// 		Claims: ,

// 	}))
// }
