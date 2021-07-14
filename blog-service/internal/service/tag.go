package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/request"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

func (s *Service) CountTag(param *request.CountTagRequest) (int, error) {
	return s.dao.CountTag(param.Name, param.State)
}

func (s *Service) GetTagList(param *request.TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (s *Service) CreateTag(param *request.CreateTagRequest) error {
	return s.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (s *Service) UpdateTag(param *request.UpdateTagRequest) error {
	return s.dao.UpdateTag(param.ID, param.Name, param.State, param.UpdatedBy)
}

func (s *Service) DeleteTag(param *request.DeleteTagRequest) error {
	return s.dao.DeleteTag(param.ID)
}
