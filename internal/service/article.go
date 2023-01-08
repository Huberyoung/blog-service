package service

import (
	"service/internal/model"
	"service/pkg/app"
)

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleGetRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"min=10,max=100"`
	Content       string `form:"content" binding:"min=10,max=65535"`
	CoverImageUrl string `form:"cover_image_url" binding:"min=5,max=2000"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=100"`
	Content       string `form:"content" binding:"max=65535"`
	CoverImageUrl string `form:"cover_image_url"`
	State         uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (s *Service) CountArticle(request *CountArticleRequest) (int, error) {
	return s.dao.CountArticle(request.Title, request.State)
}

func (s *Service) GetArticle(request *ArticleGetRequest) (model.Article, error) {
	return s.dao.GetArticle(request.ID)
}

func (s *Service) GetArticleList(request *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return s.dao.GetArticleList(request.Title, request.State, pager.Page, pager.PageSize)
}

func (s *Service) CreateArticle(request *CreateArticleRequest) error {
	return s.dao.CreateArticle(
		request.Title,
		request.Desc,
		request.Content,
		request.CoverImageUrl,
		request.State,
		request.CreatedBy,
	)
}

func (s *Service) UpdateArticle(request *UpdateArticleRequest) error {
	return s.dao.UpdateArticle(
		request.ID,
		request.Title,
		request.Desc,
		request.Content,
		request.CoverImageUrl,
		request.State,
		request.ModifiedBy,
	)
}
func (s *Service) DeleteArticle(request *DeleteArticleRequest) error {
	return s.dao.DeleteArticle(request.ID)
}
