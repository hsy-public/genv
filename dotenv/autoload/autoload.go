package autoload

import "github.com/sakirsensoy/genv/dotenv"

func init() {
	// Load environment variables from .env file
	// 读取当前跟路径下的.env文件内容，PROJECT_ENV_FILE，默认是.env

	dotenv.Load()
}
