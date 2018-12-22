package config

import (
	"time"
	"github.com/kelseyhightower/envconfig"
)

// Configuration for this example app
type Configuration struct {
	ServicePort int `envconfig:"SERVICE_PORT" required:"false" default:"8082"`

	// DB config
	DBHost         string        `envconfig:"DB_HOST" required:"false" default:"127.0.0.1"`
	DBUser         string        `envconfig:"DB_USER" required:"false" default:"root"`
	DBPass         string        `envconfig:"DB_PASS" required:"false" default:"root" redact:"true"`
	DBPort         int           `envconfig:"DB_PORT" required:"false" default:"3306"`
	DBName         string        `envconfig:"DB_PORT" required:"false" default:"drop_rules"`
	DBReadTimeout  time.Duration `envconfig:"DB_READ_TIMEOUT" required:"false" default:"10s"`
	DBWriteTimeout time.Duration `envconfig:"DB_WRITE_TIMEOUT" required:"false" default:"10s"`
}

// LoadConfig loads environment variables with the prefix
func LoadConfig() (Configuration, error) {
	cfg := Configuration{}
	err := envconfig.Process("BOUNCECM", &cfg)
	return cfg, err
}