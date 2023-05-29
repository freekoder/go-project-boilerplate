package config

type Config struct {
	Logging LoggingConfig `yaml:"logging"`
}

type LoggingConfig struct {
	Level   string `yaml:"level" env:"DEBUG_LEVEL" env-default:"DEBUG"`
	LogFile string `yaml:"log-file" env:"LOG_FILE" env-default:"service.log"`
}
