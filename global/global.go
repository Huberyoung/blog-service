package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerS
	AppSetting      *setting.AppS
	DataBaseSetting *setting.DatabaseS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
	JwtSetting      *setting.JwtS
)
