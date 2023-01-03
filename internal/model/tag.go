package model

import "service/pkg/app"

type TagSwagger struct {
	List  []*Tag
	Paper *app.Pager
}
