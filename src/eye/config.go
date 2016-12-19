package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

var (
	Setting = new(Config)
)

type Config struct {
	CpuNum int `yaml:"CpuNum"`
	Worker int `yaml:"Worker"`

	Log struct {
		LogFile  string `yaml:"LogFile"`
		LogLevel int    `yaml:"LogLevel"`
	} `yaml:"Log"`

	LocalAddr   string `yaml:"LocalAddr"`
	RedisAddr   string `yaml:"RedisAddr"`
	RedisPasswd string `yaml:"RedisPasswd"`
}

func LoadConfig() error {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./eye config")
		os.Exit(-1)
	}
	// 加载配置
	filePath := os.Args[1]

	configuration, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Load config file failed, %v", err)
		os.Exit(-2)
	} else {
		err := yaml.Unmarshal(configuration, &Setting)
		if err != nil {
			fmt.Printf("Unmarshal config file failed, %v", err)
			return err
		}
	}

	if Setting.CpuNum < 1 {
		Setting.CpuNum = 1
	}
	if Setting.Worker < 1 {
		Setting.Worker = 1
	}
	return nil
}
