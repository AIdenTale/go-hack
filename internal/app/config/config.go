package config

import (
	"github.com/spf13/viper"
)

// PostgresConfig содержит параметры подключения к PostgreSQL.
// Используется для инициализации пула соединений.
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// EchoConfig содержит параметры запуска HTTP-сервера Echo.
type EchoConfig struct {
	Address string
}

// LoggerConfig содержит параметры логгера.
type LoggerConfig struct {
	Level string
}

// Config агрегирует все параметры конфигурации приложения.
type Config struct {
	Postgres PostgresConfig
	Echo     EchoConfig
	Logger   LoggerConfig
	ML       MLConfig
}

// LoadConfig загружает конфигурацию из YAML-файла по указанному пути.
func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
