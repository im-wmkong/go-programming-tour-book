package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/request"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

func IDValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Params.Get("id"); ok {
			param := request.UriIDRequest{}
			if valid, errs := app.BindUriAndValid(c, &param); !valid {
				global.Logger.Errorf(c, "app.BindUriAndValid errs: %v", errs)
				app.NewResponse(c).ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
				c.Abort()
			}
			c.Next()
		}
		c.Next()
	}
}
