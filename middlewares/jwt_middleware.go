package middlewares

import (
	"acp8/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GenerateJWT(userId int) (string, error) {

	// payload
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// kunci
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractJWTToUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}

func ExtractTokenUserRole(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userRole := claims["Role"].(string)
		return userRole
	}
	return ""
}
