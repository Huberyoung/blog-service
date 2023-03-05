package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	CreatedOn  uint   `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn uint   `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  uint   `json:"deleted_on"`
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

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallBack)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallBack)
	db.Callback().Delete().Replace("gorm:delete", deleteCallBack)

	db.DB().SetMaxIdleConns(ds.MaxIdleConnection)
	db.DB().SetMaxOpenConns(ds.MaxOpenConnection)
	return db, nil
}

func updateTimeStampForCreateCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createdTimeField, ok := scope.FieldByName("CreatedOn"); ok && createdTimeField.IsBlank {
			_ = createdTimeField.Set(nowTime)
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok && modifyTimeField.IsBlank {
			_ = modifyTimeField.Set(nowTime)
		}
	}
}

func updateTimeStampForUpdateCallBack(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		var sql string
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			sql = fmt.Sprintf("UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)
		} else {
			sql = fmt.Sprintf("DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)
		}
		scope.Raw(sql).Exec()
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
