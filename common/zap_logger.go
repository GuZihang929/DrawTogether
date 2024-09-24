package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const zapLogFilePrefix string = "log/"
const zapLogFileSuffix string = ".log"

var (
	ZapLog         *zap.Logger
	zapLogFileName string
	SugaredLogger  *zap.SugaredLogger
)

func InitSugaredLogger(level int8, content string) {
	Filename := zapLogFilePrefix + content + zapLogFileSuffix
	logFile := &lumberjack.Logger{
		Filename:   Filename,
		MaxSize:    100,  // 日志最大保存100M
		MaxBackups: 100,  // 旧日志保留100个备份
		MaxAge:     10,   // 最多保留10天的日志 和MaxBackups参数配置1个就可以
		LocalTime:  true, // 使用当前日期
	}

	// 配置日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 配置核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(logFile),
		zapcore.Level(level),
	)
	// 创建Logger
	ZapLog = zap.New(core, zap.AddCaller())
	SugaredLogger = ZapLog.Sugar()

}
