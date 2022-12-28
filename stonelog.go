package stonelog

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

type Options struct {
	MaxSize    int    `json:"maxsize" yaml:"maxsize"`
	MaxBackups int    `json:"maxbackups" yaml:"maxbackups"`
	MaxAge     int    `json:"maxage" yaml:"maxage"`
	Out        string `json:"out" yaml:"out"`
	FileName   string `json:"filename" yaml:"filename"`
	LogLevel   string `json:"loglevel" yaml:"loglevel"`
}

//zapcore.Core需要三个配置——Encoder(编码器)，WriteSyncer(输出器)，LogLevel(日志等级)。
func Init(ops *Options) {
	encoder := getEncoder()
	writer := getWriter(ops)
	logLevel, err := zapcore.ParseLevel(ops.LogLevel)
	if err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writer, logLevel)
	logger = zap.New(core, zap.AddCaller())
	sugar = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.TimeKey = "time"
	encoderConfig.MessageKey = "message"
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriter(ops *Options) zapcore.WriteSyncer {
	switch ops.Out {
	case "file":
		return getFileWriter(ops)
	default:
		return getConsoleWriter()
	}
}

func getFileWriter(ops *Options) zapcore.WriteSyncer {
	lumberjackLogger := lumberjack.Logger{
		Filename:   ops.FileName,
		MaxSize:    ops.MaxSize,
		MaxBackups: ops.MaxBackups,
		MaxAge:     ops.MaxAge,
		Compress:   false,
	}
	return zapcore.AddSync(&lumberjackLogger)
}

func getConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func Debugf(format string, v ...interface{}) {
	sugar.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	sugar.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	sugar.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	sugar.Errorf(format, v...)
}

func Debug(format string, fields ...zap.Field) {
	logger.Debug(format, fields...)
}

func Info(format string, fields ...zap.Field) {
	logger.Info(format, fields...)
}

func Warn(format string, fields ...zap.Field) {
	logger.Warn(format, fields...)
}

func Error(format string, fields ...zap.Field) {
	logger.Error(format, fields...)
}
