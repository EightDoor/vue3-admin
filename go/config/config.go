package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type DBConfig struct {
	Host string `json:"host"`
	Port int64	`json:"port"`
	User string `json:"user"`
	Passwd string `json:"passwd"`
	Database string `json:"database"`
}

var (
	DbConfig DBConfig
)

func ParserConfig()  {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	config := cfg.Section("DB")
	DbConfig.Host = config.Key("host").String()
	DbConfig.Database = config.Key("database").String()
	DbConfig.Port, _ = config.Key("port").Int64()
	DbConfig.User = config.Key("user").String()
	DbConfig.Passwd = config.Key("passwd").String()
}