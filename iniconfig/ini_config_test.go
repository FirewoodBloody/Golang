package iniconfig

import (
	"fmt"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}
type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

func TestIniconfig(t *testing.T) {
	var conf Config
	err := UnMarshalFile("./config.ini", &conf)
	if err != nil {
		t.Errorf("test filed:%v", err)
	}
	fmt.Printf("%#v", conf)

	err = MarshalFile(conf, "D:/ini.conf")
}
