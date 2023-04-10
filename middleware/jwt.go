package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"log"
)

var (
	identityKey    = "id"
	AuthMiddleware *jwt.HertzJWTMiddleware
)

type login struct {
	Username string `form:"username,required" json:"username,required"`
	Password string `form:"password,required" json:"password,required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func PingHandler(c context.Context, ctx *app.RequestContext) {
	user, _ := ctx.Get(identityKey)
	ctx.JSON(200, utils.H{
		"message": fmt.Sprintf("username:%v", user.(*User).UserName),
	})
}

func initJwtMiddleWare() {
	AuthMiddleware = &jwt.HertzJWTMiddleware{
		Realm:            "",
		SigningAlgorithm: "",
		Key:              nil,
		KeyFunc:          nil,
		Timeout:          0,
		MaxRefresh:       0,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginValues login
			if err := c.BindAndValidate(&loginValues); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginValues.Username
			password := loginValues.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Hertz",
					FirstName: "CloudWeGo",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse:   nil,
		LogoutResponse:  nil,
		RefreshResponse: nil,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		IdentityKey:                 "",
		TokenLookup:                 "",
		TokenHeadName:               "",
		WithoutDefaultTokenHeadName: false,
		TimeFunc:                    nil,
		HTTPStatusMessageFunc:       nil,
		PrivKeyFile:                 "",
		PrivKeyBytes:                nil,
		PubKeyFile:                  "",
		PrivateKeyPassphrase:        "",
		PubKeyBytes:                 nil,
		SendCookie:                  false,
		CookieMaxAge:                0,
		SecureCookie:                false,
		CookieHTTPOnly:              false,
		CookieDomain:                "",
		SendAuthorization:           false,
		DisabledAbort:               false,
		CookieName:                  "",
		CookieSameSite:              0,
		ParseOptions:                nil,
	}
	if err := AuthMiddleware.MiddlewareInit(); err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

}

func main() {
	h := server.Default()

	h.POST("/login", AuthMiddleware.LoginHandler)

	h.NoRoute(AuthMiddleware.MiddlewareFunc(), func(ctx context.Context, c *app.RequestContext) {
		claims := jwt.ExtractClaims(ctx, c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, map[string]string{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := h.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", AuthMiddleware.RefreshHandler)
	auth.Use(AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/ping", PingHandler)
	}

	h.Spin()
}
