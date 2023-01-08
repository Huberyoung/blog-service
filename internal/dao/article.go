package dao

import (
	"service/internal/model"
	"service/pkg/app"
)

func (d *Dao) CountArticle(title string, state uint8) (int, error) {
	article := model.Article{Title: title, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticle(id uint32) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Get(d.engine, id)
}

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(title, desc, content, imgUrl string, state uint8, createdBy string) error {
	m := &model.Model{CreatedBy: createdBy}
	article := model.Article{
		Model:         m,
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: imgUrl,
		State:         state,
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title, desc, content, imgUrl string, state uint8, modifiedBy string) error {
	Article := model.Article{Model: &model.Model{ID: id}}
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

	if imgUrl != "" {
		values["cover_image_url"] = imgUrl
	}
	return Article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}
