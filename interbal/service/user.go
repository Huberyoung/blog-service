package service

import (
	"blog-service/interbal/model"
	"errors"
)

type GetUserRequest struct {
	Username string `form:"username" binding:"required,min=1,max=100"`
	Password string `form:"password" binding:"required,min=1,max=100"`
}

func (srv *Service) GetUser(request *GetUserRequest) (model.User, error) {
	return srv.dao.GetUser(request.Username, request.Password)
}

func (srv *Service) CheckUser(request *GetUserRequest) error {
	user, err := srv.GetUser(request)
	if err != nil {
		return err
	}

	if user.ID > 0 {
		return nil
	}
	return errors.New("auto info does not exist")
}
