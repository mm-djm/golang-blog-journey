package util

import (
	"golang-blog-journey/util/log"
)

type Config struct {
	ListenIP   string        `json:"listen_ip"`
	ListenPort int           `json:"listen_port"`
	LogConfig  log.LogConfig `json:"log_config"`
	DBConfig   DBConfig      `json:"db_config"`
	Ciphers    Ciphers       `json:"ciphers"`
}

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}

type Ciphers struct {
	PwdKey string `json:"pwd_key"`
}
