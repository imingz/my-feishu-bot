package env

import (
	"log/slog"

	"github.com/spf13/pflag"
)

func init() {
	pflag.StringVarP(&Active, "env", "e", DEV, "运行环境:\ndev: 开发环境\npro: 正式环境\n")
	pflag.Parse()

	if !pflag.CommandLine.Changed("env") {
		slog.Warn("未检测到 env, 使用默认值 dev")
	}

	switch Active {
	case DEV:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case PRO:
	default:
		slog.Warn("未知的 env, 使用默认值 dev")
	}
}

const (
	DEV = "dev" // 开发环境
	PRO = "pro" // 正式环境
)

var Active string
