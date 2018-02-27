package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/model"
)

var authMiddleware jwt.GinJWTMiddleware

// Auth middleware
func Auth() *jwt.GinJWTMiddleware {
	return &authMiddleware
}

func init() {
	authMiddleware = jwt.GinJWTMiddleware{
		Realm:      "gin skeleton",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			_, err := model.LoginByEmailAndPassword(userId, password)
			if err != nil {
				return userId, false
			}
			return userId, true
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" || userId == "hyper" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}
