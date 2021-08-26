package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"go.etcd.io/etcd/clientv3"
	"gopkg.in/yaml.v2"
	"strings"
	"time"
)

type Config struct {
	System     System
	ETCDConfig clientv3.Config
	duration   time.Duration

	name string
}

func (c *Config) SetSystem(system System) *Config {
	c.System = system
	return c
}

func (c *Config) SetConfig(config clientv3.Config) *Config {
	c.ETCDConfig = config
	return c
}

func (c *Config) SetContextTimeout(duration time.Duration) *Config {
	c.duration = duration
	return c
}

//Get 从etcd中获取,根据key
func (c *Config) Get(key string) (*clientv3.GetResponse, error) {
	cli, err := clientv3.New(c.ETCDConfig)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cli.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), c.duration)
	defer cancel()

	response, err := cli.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//Set 设置值
func (c *Config) Set(key string, value string) error {
	cli, err := clientv3.New(c.ETCDConfig)
	if err != nil {
		return err
	}
	defer func() {
		_ = cli.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), c.duration)
	defer cancel()

	_, err = cli.Put(ctx, key, value)
	if err != nil {
		return err
	}

	return nil
}

//SetConfigName 获取配置项目Key
func (c *Config) SetConfigName(params ...string) {
	c.name = "/" + strings.Join(params, "/")
}

//GetByName 根据名称获取配置信息
func (c *Config) GetByName() (*clientv3.GetResponse, error) {
	return c.Get(c.name)
}

//Unmarshal 反序列化
func (c *Config) Unmarshal(data []byte, ob interface{}) error {
	fileType := strings.ToLower(c.System.FileType())
	if c.System.FileType() == "yaml" {
		return yaml.Unmarshal(data, ob)
	}
	if fileType == "json" {
		return json.Unmarshal(data, ob)
	}
	return errors.New("文件类型不支持")
}
