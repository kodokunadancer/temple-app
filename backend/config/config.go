package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string

	LogFile string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("../config.ini")
	if err != nil {
		log.Fatal(err)
	}

	Config = ConfigList{
		User:     cfg.Section("db").Key("user").String(),
		Password: cfg.Section("db").Key("password").String(),
		Host:     cfg.Section("db").Key("host").String(),
		Port:     cfg.Section("db").Key("port").String(),
		DbName:   cfg.Section("db").Key("dbName").String(),
		LogFile:  cfg.Section("log").Key("log_file").String(),
	}
}
