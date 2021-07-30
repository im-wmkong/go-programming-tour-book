package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/request"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) GetAuth(c *gin.Context) {
	param := request.AuthRequest{}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c, &param); !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.CheckAuth(&param); err != nil {
		global.Logger.Errorf(c, "service.CheckAuth err: %v", err)
		response.ToErrResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
