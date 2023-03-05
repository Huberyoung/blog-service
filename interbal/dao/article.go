package dao

import (
	"blog-service/interbal/model"
	"blog-service/pkg/app"
)

func (d *Dao) GetArticle(id uint) (model.Article, error) {
	Article := &model.Article{Model: &model.Model{ID: id}}
	return Article.Get(d.engine)
}

func (d *Dao) CountArticle(title string, state uint8) (int, error) {
	Article := &model.Article{Title: title, State: state}
	return Article.Count(d.engine)
}

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	Article := &model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return Article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(title string, state uint8, createdBy string) error {
	Article := &model.Article{Model: &model.Model{CreatedBy: createdBy}, Title: title, State: state}
	return Article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint, state uint8, title, desc, content, coverImageUrl, modifiedBy string) error {
	Article := &model.Article{
		Model: &model.Model{ID: id},
	}

	values := map[string]any{
		"state":       state,
		"modified_by": modifiedBy,
	}

	if title != "" {
		values["title"] = title
	}

	if desc != "" {
		values["desc"] = desc
	}

	if content != "" {
		values["content"] = content
	}

	if coverImageUrl != "" {
		values["cover_image_url"] = coverImageUrl
	}

	return Article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint) error {
	Article := &model.Article{Model: &model.Model{ID: id}}
	return Article.Delete(d.engine)
}
