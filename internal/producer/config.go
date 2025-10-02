package producer

// Config описывает структуру producer.yaml
// для генерации и отправки данных в bpm и trac.
type Config struct {
	BPM  EndpointConfig `mapstructure:"bpm"`
	Trac EndpointConfig `mapstructure:"trac"`
}

type EndpointConfig struct {
	Mean     float64 `mapstructure:"mean"`
	Count    int     `mapstructure:"count"`
	Endpoint string  `mapstructure:"endpoint"`
	FreqHz   float64 `mapstructure:"freq_hz"`
}
