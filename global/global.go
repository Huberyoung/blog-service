package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
)

var (
	ServerSetting   *setting.ServerS
	AppSetting      *setting.AppS
	DataBaseSetting *setting.DatabaseS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
	JwtSetting      *setting.JwtS
	EmailSetting    *setting.EmailS
	Tracer          opentracing.Tracer
)
