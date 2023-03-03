package service

import (
	"blog-service/global"
	"blog-service/interbal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(tx context.Context) Service {
	service := Service{ctx: tx}
	service.dao = dao.New(global.DBEngine)
	return service
}
