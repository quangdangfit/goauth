package config

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/rinard84/gokit/library/conf"
	"github.com/rinard84/gokit/library/database"
	"github.com/spf13/viper"
)

const (
	DevelopmentEnvironment = "dev"
	ProductionEnvironment  = "prod"
)

// Config application
type Config struct {
	conf.Base     `mapstructure:",squash"`
	ServiceName   string               `mapstructure:"service_name"`
	MySQL         database.MySQLConfig `mapstructure:"mysql"`
	Redis         Redis                `mapstructure:"redis"`
	Kafka         Kafka                `mapstructure:"kafka"`
	Elasticsearch ElasticsearchConfig  `mapstructure:"elasticsearch"`
	Jwt           Jwt                  `mapstructure:"jwt"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Kafka struct {
	Broker    string               `mapstructure:"broker"`
	Producers map[string]KafkaItem `mapstructure:"producers"`
	Consumers map[string]KafkaItem `mapstructure:"consumers"`
}

type KafkaItem struct {
	Topic           string `mapstructure:"topic"`
	GroupId         string `mapstructure:"group_id"`
	SessionTimeout  int    `mapstructure:"session_timeout"`
	OffsetResetType string `mapstructure:"offset_reset_type"`
	OffsetStore     bool   `mapstructure:"offset_store"`
}

type MongoConfig struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

type ElasticsearchConfig struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Jwt struct {
	AccessTokenPrivateKey string `mapstructure:"access_token_private_key"`
	AccessTokenPublicKey  string `mapstructure:"access_token_public_key"`
}

func loadDefaultConfig() *Config {
	c := &Config{
		Base:  *conf.DefaultBaseConfig(),
		MySQL: database.MySQLDefaultConfig(),
	}

	return c
}

// Load system env config
func Load() (*Config, error) {
	/**
	|-------------------------------------------------------------------------
	| hacking to load reflect structure config into env
	|-----------------------------------------------------------------------*/
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	/**
	|-------------------------------------------------------------------------
	| You should set default config value here
	| 1. Populate the default value in (Source code)
	| 2. Then merge from config (YAML) and OS environment
	|-----------------------------------------------------------------------*/
	c := loadDefaultConfig()
	if configBuffer, err := json.Marshal(c); err != nil {
		log.Println("Oops! Marshal config is failed. ", err)
		return nil, err
	} else if err := viper.ReadConfig(bytes.NewBuffer(configBuffer)); err != nil {
		log.Println("Oops! Read default config is failed. ", err)
		return nil, err
	}
	if err := viper.MergeInConfig(); err != nil {
		log.Println("Read config file failed.", err)
	}
	// Populate all config again
	err := viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
