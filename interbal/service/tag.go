package service

import (
	"blog-service/interbal/model"
	"blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"min=3,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"min=3,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	Id         uint   `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	Id uint `form:"id" binding:"required,gte=1"`
}

func (srv *Service) CountTag(request *CountTagRequest) (int, error) {
	return srv.dao.CountTag(request.Name, request.State)
}

func (srv *Service) ListTag(request *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return srv.dao.GetTagList(request.Name, request.State, pager.Page, pager.PageSize)
}

func (srv *Service) CreateTag(request *CreateTagRequest) error {
	return srv.dao.CreateTag(request.Name, request.State, request.CreatedBy)
}

func (srv *Service) UpdateTag(request *UpdateTagRequest) error {
	return srv.dao.UpdateTag(request.Id, request.Name, request.State, request.ModifiedBy)
}

func (srv *Service) DeleteTag(request *DeleteTagRequest) error {
	return srv.dao.DeleteTag(request.Id)
}
