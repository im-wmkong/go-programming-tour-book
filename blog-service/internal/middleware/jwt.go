package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			code  = errcode.Success
		)

		if s, exists := c.GetQuery("token"); exists {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			code = errcode.InvalidParams
		} else {
			if _, err := app.ParseToken(token); err != nil {
				switch err.(jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errcode.UnauthorizedTokenTimeout
				default:
					code = errcode.UnauthorizedTokenError
				}
			}
		}

		if code != errcode.Success {
			app.NewResponse(c).ToErrResponse(code)
			c.Abort()
		}
		c.Next()
	}
}
