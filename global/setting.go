package global

import (
	"github.com/jinzhu/gorm"
	"service/pkg/logger"
	"service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
