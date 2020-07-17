package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wsloong/blog-service/global"
	"github.com/wsloong/blog-service/pkg/setting"
)

type Model struct {
	ID         uint32    `grom:"primary_key" json:"id"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	DeletedOn  time.Time `json:"deleted_on"`
	IsDel      bool      `json:"is_del"`
}

func NewDBEngine(dbSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(dbSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			dbSetting.UserName,
			dbSetting.Password,
			dbSetting.Host,
			dbSetting.DBName,
			dbSetting.Charset,
			dbSetting.ParseTime))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	// 替换回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm.update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)
	return db, nil
}

// 新增行为的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreateOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm.delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deleteOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDeleField, hasIsDelField := scope.FieldByName("IsDel")

		// 判断是否存在 DeleteOn 和 IsDel 字段
		// 如果存在执行 UPDATE 操作进行软删除
		// 否则执行硬删除
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v, %v=%v%v%v",
				scope.QuotedTableName(), // 获取表名
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDeleField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else { // 执行硬删除
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
