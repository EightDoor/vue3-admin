package middleware

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"zhoukai/configure"
)

// 设置日志级别、输出格式和日志文件的路径
func SetLogs(logLevel zapcore.Level, logFormat, fileName string) {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        configure.LEVLE_KEY,
		LevelKey:       configure.LEVLE_KEY,
		NameKey:        configure.NAME_KEY,
		CallerKey:     configure.CALLER_KEY,
		MessageKey:     configure.MESSAGE_KEY,
		StacktraceKey:  configure.STACKTRACE_KEY,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 大写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器(相对路径+行号)
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志输出格式
	var encoder zapcore.Encoder
	switch logFormat {
	case configure.LOGFORMAT_JSON:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 添加日志切割归档功能
	hook := lumberjack.Logger{
		Filename:   fileName,    // 日志文件路径
		MaxSize:    configure.MAX_SIZE,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: configure.MAX_BACKUPS, // 日志文件最多保存多少个备份
		MaxAge:     configure.MAX_AGE,     // 文件最多保存多少天
		Compress:   true,        // 是否压缩
	}

	core := zapcore.NewCore(
		encoder, // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(logLevel), // 日志级别
	)

	// 开启文件及行号
	caller := zap.AddCaller()
	// 开启开发模式，堆栈跟踪
	development := zap.Development()
	// 构造日志
	logger := zap.New(core, caller, development)

	// 将自定义的logger替换为全局的logger
	zap.ReplaceGlobals(logger)
}