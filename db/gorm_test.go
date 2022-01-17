package db

import (
	"fmt"
	"github.com/fatih/structs"
	"testing"

	"github.com/icarus-go/component/db/config"
	"go.uber.org/zap"
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

type A struct {
	Foo  string
	Test string
}

func (a *A) PrintFoo() {
	fmt.Println("Foo value is " + a.Foo)
}

func Test_get_FieldName(t *testing.T) {
	names := structs.Names(A{})
	fmt.Println(names) // ["Foo", "Bar"]
}
