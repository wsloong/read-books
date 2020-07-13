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
	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)
	return db, nil
}
