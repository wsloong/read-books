package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/wsloong/blog-service/global"
	"github.com/wsloong/blog-service/pkg/util"
)

// FileType 为int的类型别名
type FileType int

const TypeImage FileType = iota + 1

// 获取文件名称，并进行md5加密之后返回
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 获取文件保存地址，直接返回配置，以便后续有所调整
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// 检查是否是允许的文件后缀
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
				return true
			}
		}
	}
	return false
}

// 检查文件大小是否超出限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// 创建保存上传文件的目录
func CreateSavePath(dst string, perm os.FileMode) error {
	return os.MkdirAll(dst, perm)
}

// 保存上传文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
