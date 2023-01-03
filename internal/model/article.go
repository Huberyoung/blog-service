package model

import "service/pkg/app"

type ArticleSwagger struct {
	List  []*Article
	Paper *app.Pager
}
