package customload

import (
	"github.com/sakirsensoy/genv/dotenv"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	// 获取当前文件所在路径
	var envFilePath string
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	// 获取项目根路径
	rootDir, err := filepath.Abs(filepath.Join(dir, ".."))
	if err != nil {
		panic(err)
	}

	// 检查genv.env文件是否存在，如果不存在则默认使用.env文件
	genvFilePath := filepath.Join(rootDir, "genv.env")
	if _, err := os.Stat(genvFilePath); err == nil {
		// 解析genv.env文件并获取PROJECT_ENV_NAME字段的值
		dotenv.Load(genvFilePath) // 使用正确的方法进行解析
		envFilePath = os.Getenv("PROJECT_ENV_NAME")
		if envFilePath == "" {
			envFilePath = ".env"
		}
	}

	// 加载指定的文件
	dotenv.Load(envFilePath)
}
