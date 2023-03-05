package service

import (
	"blog-service/interbal/model"
	"blog-service/pkg/app"
)

type GetArticleRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type ListArticleRequest struct {
	Title string `form:"title"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=1,max=10"`
	Desc          string `form:"desc" binding:"required,min=3,max=100"`
	Content       string `form:"content" binding:"required,min=3,max=10000"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy     string `form:"created_by" binding:"required,min=1,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint   `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=10"`
	Desc          string `form:"desc" binding:"max=100"`
	Content       string `form:"content" binding:"max=10000"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=1,max=100"`
}

type DeleteArticleRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

func (srv *Service) GetArticle(request *GetArticleRequest) (model.Article, error) {
	return srv.dao.GetArticle(request.ID)
}

func (srv *Service) CountArticle(request *ListArticleRequest) (int, error) {
	return srv.dao.CountArticle(request.Title, request.State)
}

func (srv *Service) ListArticle(request *ListArticleRequest, pager *app.Pager) ([]*model.Article, error) {
	return srv.dao.GetArticleList(request.Title, request.State, pager.Page, pager.PageSize)
}

func (srv *Service) CreateArticle(request *CreateArticleRequest) error {
	return srv.dao.CreateArticle(request.Title, request.Desc, request.Content, request.CoverImageUrl, request.State, request.CreatedBy)
}

func (srv *Service) UpdateArticle(request *UpdateArticleRequest) error {
	return srv.dao.UpdateArticle(request.ID, request.State, request.Title, request.Desc, request.Content, request.CoverImageUrl, request.ModifiedBy)
}

func (srv *Service) DeleteArticle(request *DeleteArticleRequest) error {
	return srv.dao.DeleteArticle(request.ID)
}
