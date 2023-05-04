package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"financial_management/consts"
	"financial_management/setting"
	"financial_management/util"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			util.Response(c, consts.AuthErrorCode, nil)
			c.Abort()
			return
		}
		claim := ParseToken(strings.TrimPrefix(tokenString, "Bearer "))
		if claim == nil {
			util.Response(c, consts.AuthErrorCode, nil)
			c.Abort()
			return
		}
		c.Set("user_id", claim.UserId)
		return
	}
}

var (
	tokenExpireDuration = time.Hour * 24 * 7
)

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func GenerateToken(userID int) string {
	newClaims := Claims{
		userID, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(),
			Issuer:    "flying",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, _ := token.SignedString([]byte(setting.Config.JwtKey))
	return tokenString
}

func ParseToken(token string) *Claims {
	var newClaims = new(Claims)
	tmpToken, err := jwt.ParseWithClaims(token, newClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.Config.JwtKey), nil
	})
	if err != nil {
		return nil
	}
	if tmpToken != nil {
		if tmpClaims, ok := tmpToken.Claims.(*Claims); ok && tmpToken.Valid {
			return tmpClaims
		}
	}
	return nil
}
