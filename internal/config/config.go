package config

import (
	"os"
)

const (
	DEVELOPER    = "developer"
	HOMOLOGATION = "homologation"
	PRODUCTION   = "production"
)

type Config struct {
	PORT string `json:"port"`
	Mode string `json:"mode"`

	*MySQLConfig
}

type MySQLConfig struct {
	DB_DRIVE                  string `json:"db_drive"`
	DB_HOST                   string `json:"db_host"`
	DB_PORT                   string `json:"db_port"`
	DB_USER                   string `json:"db_user"`
	DB_PASS                   string `json:"db_pass"`
	DB_NAME                   string `json:"db_name"`
	DB_DSN                    string `json:"-"`
	DB_SET_MAX_OPEN_CONNS     int    `json:"db_set_max_open_conns"`
	DB_SET_MAX_IDLE_CONNS     int    `json:"db_set_max_idle_conns"`
	DB_SET_CONN_MAX_LIFE_TIME int    `json:"db_set_conn_max_life_time"`
	SRV_DB_SSL_MODE           bool   `json:"srv_db_ssl_mode"`
}

func NewConfig() *Config {
	conf := defaultConf()

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.PORT = SRV_PORT
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST != "" {
		conf.DB_HOST = DB_HOST
	}

	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT != "" {
		conf.DB_PORT = DB_PORT
	}

	DB_USER := os.Getenv("DB_USER")
	if DB_USER != "" {
		conf.DB_USER = DB_USER
	}

	DB_PASS := os.Getenv("DB_PASS")
	if DB_PASS != "" {
		conf.DB_PASS = DB_PASS
	}

	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME != "" {
		conf.DB_NAME = DB_NAME
	}

	return conf
}

func defaultConf() *Config {
	default_conf := Config{
		PORT: "8080",
		Mode: DEVELOPER,
		MySQLConfig: &MySQLConfig{
			DB_DRIVE: "mysql",
			DB_HOST:  "localhost",
			DB_PORT:  "3306",
			DB_USER:  "root",
			DB_PASS:  "supersenha",
			DB_NAME:  "eulabs_db_dev",
		},
	}

	return &default_conf
}
