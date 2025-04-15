package database

import (
	"fmt"
)

// DBConfig used to set config for database.
type DBConfig interface {
	String() string
	DSN() string
}

// Config used to set base config for all database types.
type Config struct {
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	Database string `json:"database" mapstructure:"database" yaml:"database"`
	Port     int    `json:"port" mapstructure:"port" yaml:"port"`
	Username string `json:"username" mapstructure:"username" yaml:"username"`
	Password string `json:"password" mapstructure:"password" yaml:"password"`
	Protocol string `json:"protocol" mapstructure:"protocol" yaml:"protocol"`
	Options  string `json:"options" mapstructure:"options" yaml:"options"`
}

// DSN returns the Domain Source Name.
func (c Config) DSN() string {
	options := c.Options
	if options != "" {
		if options[0] != '?' {
			options = "?" + options
		}
	}
	if c.Protocol == "" {
		c.Protocol = "tcp"
	}
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s%s",
		c.Username,
		c.Password,
		c.Protocol,
		c.Host,
		c.Port,
		c.Database,
		options)
}

func (c PostgreSQLConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s%s", c.Username, c.Password, c.Host, c.Port, c.Database, c.Options)
}

// MySQLConfig used to set config for MySQL.
type MySQLConfig struct {
	Config `mapstructure:",squash"`
}

// String returns MySQL connection URI.
func (c MySQLConfig) String() string {
	return fmt.Sprintf("mysql://%s", c.DSN())
}

// PostgreSQLConfig used to set config for Postgres.
type PostgreSQLConfig struct {
	Config `mapstructure:",squash"`
}

// String returns Postgres connection URI.
func (c PostgreSQLConfig) String() string {
	return fmt.Sprintf("postgresql://%s", c.DSN())
}

// PostgresSQLDefaultConfig returns default config for mysql, usually use on development.
func PostgresSQLDefaultConfig() PostgreSQLConfig {
	return PostgreSQLConfig{Config{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "sample",
		Username: "default",
		Password: "secret",
		Options:  "",
	}}
}

// MySQLDefaultConfig returns default config for mysql, usually use on development.
func MySQLDefaultConfig() MySQLConfig {
	return MySQLConfig{Config{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "sample",
		Username: "default",
		Password: "secret",
		Options:  "?parseTime=true",
	}}
}

type ClickhouseConfig struct {
	Config `mapstructure:",squash"`
}

func (c ClickhouseConfig) DSN() string {
	// clickhouse://username:password@host1:9000/database?dial_timeout=200ms&max_execution_time=60
	if c.Options == "" {
		c.Options = "dial_timeout=200ms&max_execution_time=60"
	}
	return fmt.Sprintf("%s:%s@%s:%d/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Database, c.Options)
}

// String returns Postgres connection URI.
func (c ClickhouseConfig) String() string {
	return fmt.Sprintf("clickhouse://%s", c.DSN())
}

// PostgresSQLDefaultConfig returns default config for mysql, usually use on development.
func ClickhouseDefaultConfig() ClickhouseConfig {
	return ClickhouseConfig{Config{
		Host:     "127.0.0.1",
		Port:     8123,
		Database: "default",
		Username: "default",
		Password: "secret",
		Options:  "",
	}}
}
