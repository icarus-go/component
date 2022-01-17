package log

import (
	"log"
	"os"
	"strings"
	"time"

	thisConfig "github.com/icarus-go/component/db/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Set(gormConfig *gorm.Config, params thisConfig.Params) *gorm.Config {
	mode := strings.ToLower(params.LogMode)

	Default := NewLogger(log.New(os.Stdout, "\r\n", log.LstdFlags), config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
		LogZap:        params.LogZap,
	})

	switch mode {
	case "silent":
		gormConfig.Logger = Default.LogMode(logger.Silent)
	case "error":
		gormConfig.Logger = Default.LogMode(logger.Error)
	case "warn":
		gormConfig.Logger = Default.LogMode(logger.Warn)
	case "info":
		gormConfig.Logger = Default.LogMode(logger.Info)
	default:
		gormConfig.Logger = Default.LogMode(logger.Info)
	}

	return gormConfig
}
