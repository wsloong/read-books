package global

import (
	"github.com/wsloong/blog-service/pkg/logger"
	"github.com/wsloong/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSetting
	Logger          *logger.Logger
)
