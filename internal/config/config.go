package config

type Config struct {
	Logging LoggingConfig `yaml:"logging"`
	Web     WebConfig     `yaml:"web"`
}

type LoggingConfig struct {
	Level   string `yaml:"level" env:"DEBUG_LEVEL" env-default:"DEBUG"`
	LogFile string `yaml:"log-file" env:"LOG_FILE" env-default:"service.log"`
}

type WebConfig struct {
	Port int `yaml:"port" env:"WEB_PORT" env-default:"9090"`
}
