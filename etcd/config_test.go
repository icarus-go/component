package etcd

import (
	"log"
	"testing"
)

func TestConfig_Get(t *testing.T) {
	config, err := NewGinModeConfig()
	if err != nil {
		log.Println(err.Error())
		return
	}

	config.System.FileType()

	get, err := config.Get("/tenyum-system/develop")
	if err != nil {
		log.Println(err.Error())
		return
	}

	for key, value := range get.Kvs {
		log.Println("key = ", key, "value = ", value)
	}
}
