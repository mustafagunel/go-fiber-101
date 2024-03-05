package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		Http     `yaml:"http"`
		Log      `yaml:"logger"`
		Mysql_Db `yaml:"mysqlDb"`
	}

	// Mysql_Db -.
	Mysql_Db struct {
		Host          string `env-required:"true" yaml:"host"`
		Port          string `env-required:"true" yaml:"port"`
		User          string `env-required:"true" yaml:"user"`
		Password      string `env-required:"true" yaml:"password"`
		Dbname        string `env-required:"true" yaml:"dbName"`
		MigrationPath string `env-required:"true" yaml:"migrationPath"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		Prod    *bool  `env-required:"true" yaml:"prod"    env:"APP_PROD"`
	}

	// HTTP -.
	Http struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		// MaxHeaderBytes string   `env-required:"true" yaml:"maxHeaderBytes" env:"MHB"`
		ReadTimeout      string   `env-required:"true" yaml:"readTimeout" env:"RTO"`
		WriteTimeout     string   `env-required:"true" yaml:"writeTimeout" env:"WTO"`
		AllowedOrigins   []string `env-required:"true" yaml:"allowedOrigins"`
		AllowedMethods   []string `env-required:"true" yaml:"allowedMethods"`
		AllowedHeaders   []string `env-required:"true" yaml:"allowedHeaders"`
		ExposedHeaders   []string `env-required:"true" yaml:"exposedHeaders"`
		MaxAge           int      `env-required:"true" yaml:"maxAge"`
		AllowCredentials bool     `env-required:"true" yaml:"allowCredentials"`
	}
	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

var _cfg *Config

// NewConfig returns app config.
func NewConfig(path ...string) (*Config, error) {
	cfg := &Config{}
	var err error
	if len(path) > 0 {
		err = cleanenv.ReadConfig(path[0], cfg)

		if err != nil {
			return nil, fmt.Errorf("config error: %w", err)
		}
	} else {
		//err = cleanenv.ReadConfig("../../config/config.yml", cfg)
		err = cleanenv.ReadConfig("internal/config/config_dev.yml", cfg)

		if err != nil {
			return nil, fmt.Errorf("config error: %w", err)
		}
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	_cfg = cfg
	return cfg, nil
}

func GetCfg() *Config {
	return _cfg
}
