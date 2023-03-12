package service

import (
	"blog-service/global"
	"blog-service/interbal/dao"
	"context"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(tx context.Context) Service {
	service := Service{ctx: tx}
	service.dao = dao.New(otgorm.WithContext(service.ctx, global.DBEngine))
	return service
}
