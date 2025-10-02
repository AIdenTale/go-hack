package config

// MLConfig содержит параметры для ML сервиса
type MLConfig struct {
	BaseURL string `mapstructure:"base_url"` // URL ML сервиса
	UpdateInterval int `mapstructure:"update_interval"` // Интервал обновления в секундах
}