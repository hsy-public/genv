package customload

import (
	"github.com/sakirsensoy/genv/dotenv"
	"os"
)

func init() {
	// 获取当前文件所在路径
	var envFilePath string

	// 检查genv.env文件是否存在，如果不存在则默认使用.env文件
	dotenv.Load("genv.env") // 使用正确的方法进行解析
	envFilePath = os.Getenv("PROJECT_ENV_NAME")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	// 加载指定的文件
	dotenv.Load(envFilePath)
}
