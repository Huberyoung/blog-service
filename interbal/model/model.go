package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	CreatedOn  uint   `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn uint   `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeleteOn   uint   `json:"delete_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDbEngine(ds *setting.DatabaseS) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	s := fmt.Sprintf(format, ds.Username, ds.Password, ds.Host, ds.DBName, ds.Charset, ds.ParseTime)
	db, err := gorm.Open(ds.DBType, s)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == gin.DebugMode {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(ds.MaxIdleConnection)
	db.DB().SetMaxOpenConns(ds.MaxOpenConnection)
	return db, nil
}
