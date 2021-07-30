package service

import (
	"errors"

	"github.com/go-programming-tour-book/blog-service/internal/request"
)

func (s *Service) CheckAuth(param *request.AuthRequest) error {
	auth, err := s.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exists.")
}
