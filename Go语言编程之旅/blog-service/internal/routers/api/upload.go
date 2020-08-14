package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wsloong/blog-service/global"
	"github.com/wsloong/blog-service/internal/service"
	"github.com/wsloong/blog-service/pkg/app"
	"github.com/wsloong/blog-service/pkg/convert"
	"github.com/wsloong/blog-service/pkg/errcode"
	"github.com/wsloong/blog-service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFile.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"file_access_url": fileInfo.AccessUrl})
	return
}
