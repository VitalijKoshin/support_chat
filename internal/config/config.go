package config

import (
	"supportchat/internal/apperrors"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT,required"`
	LogLevel    string `env:"LOG_LEVEL,required"`
	Port        string `env:"PORT,required"`
	MySQL       MySQLConfig
	Redis       RedisConfig
	JWT         JwtConfig
}

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST,required"`
	Port     string `env:"MYSQL_PORT,required"`
	Username string `env:"MYSQL_USERNAME,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	Database string `env:"MYSQL_DATABASE,required"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST,required"`
	Port     string `env:"REDIS_PORT,required"`
	User     string `env:"REDIS_USER,required"`
	Password string `env:"REDIS_PASS,required"`
	DB       int    `env:"REDIS_DB,required"`
}

type JwtConfig struct {
	Secret string `env:"JWT_SECRET,required"`
	Ttl    int    `env:"JWT_TTL,required"`
}

func NewConfig(envStr string) (*Config, error) {
	err := godotenv.Load(envStr)
	if err != nil {
		return nil, apperrors.EnvConfigLoadError.AppendMessage(err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return cfg, apperrors.EnvConfigParseError.AppendMessage(err)
	}

	err = parseMySQLConfig(cfg)
	if err != nil {
		return cfg, err
	}

	err = parseRedisConfig(cfg)
	if err != nil {
		return cfg, err
	}

	err = parseJwtConfig(cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func parseMySQLConfig(cfg *Config) error {
	opts := env.Options{}
	mysqlCfg := &MySQLConfig{}
	if err := env.ParseWithOptions(mysqlCfg, opts); err != nil {
		return apperrors.EnvConfigMysqlParseError.AppendMessage(err)
	}
	cfg.MySQL = *mysqlCfg
	return nil
}

func parseRedisConfig(cfg *Config) error {
	opts := env.Options{}
	redisCfg := &RedisConfig{}
	if err := env.ParseWithOptions(redisCfg, opts); err != nil {
		return apperrors.EnvConfigRedisParseError.AppendMessage(err)
	}
	cfg.Redis = *redisCfg
	return nil
}

func parseJwtConfig(cfg *Config) error {
	opts := env.Options{}
	jwtCfg := &JwtConfig{}
	if err := env.ParseWithOptions(jwtCfg, opts); err != nil {
		return apperrors.EnvConfigJwtParseError.AppendMessage(err)
	}
	cfg.JWT = *jwtCfg
	return nil
}
