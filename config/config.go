package config

import (
	"encoding/json"
	"net"
	"strings"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
	Timezone string `json:"timezone"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type Key struct {
	ACCESS_SECRET  string `json:"access_secret"`
	REFRESH_SECRET string `json:"refresh_secret"`
}

type Telegram struct {
	BOT_NAME  string `json:"bot_name"`
	BOT_TOKEN string `json:"bot_token"`
}

type General struct {
	Address  string   `json:"address"`
	Domain   string   `json:"domain"`
	Database Database `json:"database"`
	Redis    Redis    `json:"redis"`
	Key      Key      `json:"key"`
	Telegram Telegram `json:"telegram"`
}

func Parse(buf []byte) (*General, error) {
	rawCfg, err := UnmarshalRawConfig(buf)
	if err != nil {
		return nil, err
	}

	return ParseRawConfig(rawCfg)
}

func UnmarshalRawConfig(buf []byte) (*General, error) {
	rawCfg := &General{
		Address: "localhost:8080",
		Domain:  "aime.io",
		Key: Key{
			ACCESS_SECRET:  "secret",
			REFRESH_SECRET: "secret",
		},
		Database: Database{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			DbName:   "postgres",
			SSLMode:  "disable",
			Timezone: "Asia/Shanghai",
		},
		Redis: Redis{
			Host:     "localhost",
			Port:     6379,
			Password: "",
		},
		Telegram: Telegram{
			BOT_NAME:  "aime",
			BOT_TOKEN: "token",
		},
	}

	if err := json.Unmarshal(buf, rawCfg); err != nil {
		return nil, err
	}

	return rawCfg, nil
}

func ParseRawConfig(rawCfg *General) (*General, error) {
	config := &General{}

	general, err := ParseGeneral(rawCfg)
	if err != nil {
		return nil, err
	}
	config = general

	return config, nil
}

func ParseGeneral(cfg *General) (*General, error) {
	return &General{
		Address: cfg.Address,
		Domain:  cfg.Domain,
		Key: Key{
			ACCESS_SECRET:  cfg.Key.ACCESS_SECRET,
			REFRESH_SECRET: cfg.Key.REFRESH_SECRET,
		},
		Database: Database{
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			User:     cfg.Database.User,
			Password: cfg.Database.Password,
			DbName:   cfg.Database.DbName,
			SSLMode:  cfg.Database.SSLMode,
			Timezone: cfg.Database.Timezone,
		},
		Redis: Redis{
			Host:     cfg.Redis.Host,
			Port:     cfg.Redis.Port,
			Password: cfg.Redis.Password,
		},
		Telegram: Telegram{
			BOT_NAME:  cfg.Telegram.BOT_NAME,
			BOT_TOKEN: cfg.Telegram.BOT_TOKEN,
		},
	}, nil
}

func hostWithDefaultPort(host string, defPort string) (string, error) {
	if !strings.Contains(host, ":") {
		host += ":"
	}

	hostname, port, err := net.SplitHostPort(host)
	if err != nil {
		return "", err
	}

	if port == "" {
		port = defPort
	}

	return net.JoinHostPort(hostname, port), nil
}
