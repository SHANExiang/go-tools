package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go-tools/pkg/file"
	"go-tools/pkg/global"
	"go-tools/pkg/zap/internal"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := file.PathExists(global.Server.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.Server.Zap.Director)
		_ = os.Mkdir(global.Server.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.Server.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}


func InitZap() {
	global.LOG = Zap()
	zap.ReplaceGlobals(global.LOG)
}
