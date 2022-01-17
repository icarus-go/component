package zap

import (
	"fmt"
	"github.com/icarus-go/component/zap/constant"
	logs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func NewZap(config Config) *_zap {
	instance := new(_zap)
	instance.Config = config
	return instance
}

type _zap struct {
	err    error
	level  zapcore.Level
	writer zapcore.WriteSyncer
	zap    *zap.Logger
	Config
}

func (z *_zap) Initialize() {
	switch z.Level { // 初始化配置文件的Level
	case constant.Debug.Value():
		z.level = zap.DebugLevel
	case constant.Info.Value():
		z.level = zap.InfoLevel
	case constant.Warn.Value():
		z.level = zap.WarnLevel
	case constant.Error.Value():
		z.level = zap.ErrorLevel
	case constant.Dpanic.Value():
		z.level = zap.DPanicLevel
	case constant.Panic.Value():
		z.level = zap.PanicLevel
	case constant.Fatal.Value():
		z.level = zap.FatalLevel
	default:
		z.level = zap.InfoLevel
	}
	if z.writer, z.err = z.getWriteSyncer(); z.err != nil { // 使用file-rotatelogs进行日志分割
		fmt.Println(`获取WriteSyncer失败, err: `, z.err)
		return
	}
	if z.level == zap.DebugLevel || z.level == zap.ErrorLevel {
		z.zap = zap.New(z.getEncoderCore(z.writer, z.level), zap.AddStacktrace(z.level))
	} else {
		z.zap = zap.New(z.getEncoderCore(z.writer, z.level))
	}
	if z.ShowLine {
		z.zap = z.zap.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(z.zap)
}

// GetEncoderConfig 获取zapcore.EncoderConfig
func (z *_zap) GetEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "log",
		CallerKey:      "caller",
		StacktraceKey:  z.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     z.customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func (z *_zap) getEncoder() zapcore.Encoder {
	if z.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

//getEncoderCore 获取Encoder的zapcore.Core
func (z *_zap) getEncoderCore(writer zapcore.WriteSyncer, level zapcore.Level) (core zapcore.Core) {
	return zapcore.NewCore(z.getEncoder(), writer, level)
}

//customTimeEncoder 自定义日志输出时间格式
func (z *_zap) customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(z.Prefix + "2006/01/02 - 15:04:05.000"))
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: zap logger中加入file-rotatelogs
func (z *_zap) getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := logs.New(
		path.Join(z.Director, "%Y-%m-%d.log"),
		logs.WithMaxAge(7*24*time.Hour),
		logs.WithRotationTime(24*time.Hour),
	)
	if z.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}

	return zapcore.AddSync(fileWriter), err
}
