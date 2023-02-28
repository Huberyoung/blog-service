package model

type ArticleTag struct {
	*Model
	ArticleId uint `json:"article_id"`
	TagId     uint `json:"tag_id"`
}

func (A ArticleTag) TableName() string {
	return "blog_article_tag"
}
