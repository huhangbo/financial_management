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
		StuId := claim.StuId
		c.Set("stuId", StuId)
		return
	}
}

var (
	jwtKey              = []byte(setting.Config.JwtKey)
	tokenExpireDuration = time.Hour * 24 * 7
)

type Claims struct {
	StuId string
	jwt.StandardClaims
}

func GenerateToken(stuId string) string {
	newClaims := Claims{
		stuId, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(),
			Issuer:    "flying",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func ParseToken(token string) *Claims {
	var newClaims = new(Claims)
	tmpToken, err := jwt.ParseWithClaims(token, newClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
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
