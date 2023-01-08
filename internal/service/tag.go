package service

import (
	"service/internal/model"
	"service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required, gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"required, oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
}

func (s *Service) CountTag(request *CountTagRequest) (int, error) {
	return s.dao.CountTag(request.Name, request.State)
}

func (s *Service) GetTagList(request *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(request.Name, request.State, pager.Page, pager.PageSize)
}

func (s *Service) CreateTag(request *CreateTagRequest) error {
	return s.dao.CreateTag(request.Name, request.State, request.CreatedBy)
}

func (s *Service) UpdateTag(request *UpdateTagRequest) error {
	return s.dao.UpdateTag(request.ID, request.Name, request.State, request.ModifiedBy)
}

func (s *Service) DeleteTag(request *DeleteTagRequest) error {
	return s.dao.DeleteTag(request.ID)
}
