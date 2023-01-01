package model

import (
	"fmt"
	"service/global"
	"service/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeleteOn   uint32 `json:"delete_on"`
	IsDel      uint8  `json:"is_del"`
}

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"` // state 状态 0 为禁用、1 为启用
}

func (a Article) TableName() string {
	return "blog_article"
}

type BlogTag struct {
	// id ColumnKey:PRI
	Id int32 `json:"id"`
	// name 标签名称
	Name string `json:"name"`
	// created_on 创建时间
	CreatedOn int32 `json:"created_on"`
	// created_by 创建人
	CreatedBy string `json:"created_by"`
	// modified_on 修改时间
	ModifiedOn int32 `json:"modified_on"`
	// modified_by 修改人
	ModifiedBy string `json:"modified_by"`
	// deleted_on 删除时间
	DeletedOn int32 `json:"deleted_on"`
	// is_del 是否删除 0 为未删除、1 为已删除
	IsDel int8 `json:"is_del"`
	// state 状态 0 为禁用、1 为启用
	State int8 `json:"state"`
}

func (model BlogTag) TableName() string {
	return "blog_tag"
}

type ArticleTag struct {
	*Model
	ArticleId uint32 `json:"article_id"`
	TagId     uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

func NewDbEngine(d *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", d.Username, d.Password, d.Host, d.DBName, d.Charset, d.ParseTime)
	db, err := gorm.Open(d.DBType, s)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == gin.DebugMode {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(d.MaxIdleConnections)
	db.DB().SetMaxOpenConns(d.MaxOpenConnections)
	return db, nil
}
