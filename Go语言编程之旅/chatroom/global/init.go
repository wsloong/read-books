package global

import (
	"os"
	"path/filepath"
	"sync"
)

func init() {
	Init()
}

var RootDir string
var once sync.Once

func Init() {
	once.Do(func() {
		inferRootDir()
		//initConfig()
	})
}

// inferRootDir 推断出项目根目录
func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		// 这里确保项目根目录下存在template目录
		if existes(d + "/template") {
			return d
		}
		return infer(filepath.Dir(d))
	}
	RootDir = infer(cwd)
}

func existes(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
