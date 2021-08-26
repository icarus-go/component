package log

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

func GenerateConfig(mode string, logZap bool) *gorm.Config {
	mode = strings.ToLower(mode)

	_config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}

	Default := NewLogger(log.New(os.Stdout, "\r\n", log.LstdFlags), config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
		LogZap:        logZap,
	})

	switch mode {
	case "silent":
		_config.Logger = Default.LogMode(logger.Silent)
	case "error":
		_config.Logger = Default.LogMode(logger.Error)
	case "warn":
		_config.Logger = Default.LogMode(logger.Warn)
	case "info":
		_config.Logger = Default.LogMode(logger.Info)
	default:
		_config.Logger = Default.LogMode(logger.Info)
	}
	return _config
}
