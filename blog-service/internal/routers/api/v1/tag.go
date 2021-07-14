package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/request"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取单个标签
// @Produce  json
// @Router /api/v1/tags/{id} [get]
// @Param id path int true "标签 ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
func (t Tag) Get(c *gin.Context) {}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := request.TagListRequest{}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c, &param); !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&request.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "service.CountTag err: %v", err)
		response.ToErrResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "service.GetTagList err: %v", err)
		response.ToErrResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := request.CreateTagRequest{}
	response := app.NewResponse(c)

	if valid, errs := app.BindAndValid(c, &param); !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.CreateTag(&param); err != nil {
		global.Logger.Errorf(c, "service.CreateTag err: %v", err)
		response.ToErrResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := request.UpdateTagRequest{}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c, param); !valid {
		global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.UpdateTag(&param); err != nil {
		global.Logger.Errorf(c, "service.UpdateTag err: %v", err)
		response.ToErrResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := request.DeleteTagRequest{}
	response := app.NewResponse(c)
	if valid, errs := app.BindAndValid(c, &param); !valid {
		global.Logger.Errorf(c, "app.BindAndValid err: %v", errs)
		response.ToErrResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	if err := svc.DeleteTag(&param); err != nil {
		global.Logger.Errorf(c, "service.DeleteTag err: %v", err)
		response.ToErrResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
