package config

import (
	"net"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	defaultHTTPPort = "8081"
	defaultGRPCPort = "8082"
)

type databaseConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"ssl_mode"`
}

type authenticationConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type ports struct {
	HttpAddr string
	GrpcAddr string
}

type Config struct {
	Database       databaseConfig       `yaml:"database"`
	Authentication authenticationConfig `yaml:"authentication"`
	Ports          ports
}

func NewConfig(logger log.Logger) *Config {
	cfg, err := parseConfigFiles(logger, "config/db_config.yaml", "config/auth_config.yaml")
	if err != nil {
		log.With(logger, "ts_cfg", "Error parsing config files: %v", err)
		return nil
	}

	cfg.Ports = ports{
		net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort)),
		net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort)),
	}

	err = os.Setenv("AUTH_USER", cfg.Authentication.User)
	if err != nil {
		log.With(logger, "ts_cfg", "Error during setup env variable: %v", err)
		return nil
	}
	err = os.Setenv("AUTH_PASS", cfg.Authentication.Password)
	if err != nil {
		log.With(logger, "ts_cfg", "Error during setup env variable: %v", err)
		return nil
	}

	log.With(logger, "ts_cfg", "Parsed configuration")
	log.With(logger, "ts_cfg", *cfg)
	return cfg
}

func parseConfigFiles(logger log.Logger, files ...string) (*Config, error) {
	var cfg Config
	for i := 0; i < len(files); i++ {
		err := cleanenv.ReadConfig(files[i], &cfg)
		if err != nil {
			log.With(logger, "ts_cfg", "Error reading configuration from file:%v", files[i])
			return nil, err
		}
	}
	return &cfg, nil
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
