package db

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
	"pmo-test4.yz-intelligence.com/kit/component/db/config"
)

func TestNew(t *testing.T) {
	_, err := New(config.Params{}, func(instance *Gorm) {
		instance.
			SetAllowGlobalUpdate().
			SetDDLRule().
			SetAutoMigrateTables().
			SetLogger().
			SetDisableAutomaticPing().
			SetDisableNestedTransaction()
	})
	if err != nil {
		zap.L().Error("Failed to initialize")
		return
	}

}

func TestDefaultNew(t *testing.T) {
	gorm, err := DefaultNew(config.Params{})
	if err != nil {
		zap.L().Error("Failed to initialize")
		return
	}

	zap.L().Info(fmt.Sprintf("%v", gorm))

}
