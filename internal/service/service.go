package service

import (
	"context"
	"service/global"
	"service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	service := Service{ctx: ctx}
	service.dao = dao.New(global.DBEngine)
	return service
}
