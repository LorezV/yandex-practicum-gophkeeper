package server

import (
	"fmt"
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App          `yaml:"app"`
		HTTP         `yaml:"http"`
		Log          `yaml:"logger"`
		PG           `yaml:"postgres"`
		Secutiry     `yaml:"security"`
		Cache        `yaml:"cache"`
		FilesStorage `yaml:"files_storage"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" yaml:"pg_url" env:"PG_URL"`
	}
	Cache struct {
		DefaultExpiration int `yaml:"default_expiration" env:"DEFAULT_EXPIRATION"`
		CleanupInterval   int `yaml:"cleanup_interval" env:"CLEANUP_INTERVAL"`
	}
	FilesStorage struct {
		Location string `yaml:"location" env:"FILES_LOCATION"`
	}
	Secutiry struct {
		AccessTokenPrivateKey  string        `yaml:"access_token_private_key" env:"ACCESS_TOKEN_PRIVATE_KEY"`
		AccessTokenPublicKey   string        `yaml:"access_token_public_key" env:"ACCESS_TOKEN_PUBLIC_KEY"`
		RefreshTokenPrivateKey string        `yaml:"refresh_token_private_key" env:"REFRESH_TOKEN_PRIVATE_KEY"`
		RefreshTokenPublicKey  string        `yaml:"refresh_token_public_key" env:"REFRESH_TOKEN_PUBLIC_KEY"`
		AccessTokenExpiresIn   time.Duration `yaml:"access_token_expired_in" env:"ACCESS_TOKEN_EXPIRED_IN"`
		RefreshTokenExpiresIn  time.Duration `yaml:"refresh_token_expired_in" env:"REFRESH_TOKEN_EXPIRED_IN"`
		AccessTokenMaxAge      int           `yaml:"access_token_maxage" env:"ACCESS_TOKEN_MAXAGE"`
		RefreshTokenMaxAge     int           `yaml:"refresh_token_maxage" env:"REFRESH_TOKEN_MAXAGE"`
		Domain                 string        `yaml:"domain" env:"DOMAIN"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/server/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if err := cleanenv.ReadConfig(".env", cfg); err != nil {
		log.Println(err.Error())
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
