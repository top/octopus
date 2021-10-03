package config

import "fmt"

// Mysql Database configuration
type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	DbName       string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Pass         string `mapstructure:"pass" json:"pass" yaml:"pass"`
	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"`
	LogMode      bool   `mapstructure:"logMode" json:"logMode" yaml:"logMode"`
}

func (m Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", m.User, m.Pass, m.Host, m.DbName, m.Config)
}

// Cassandra Database configuration
type Cassandra struct {
	Hosts []string `mapstructure:"hosts" json:"hosts" yaml:"hosts"`
}

// Redis Database configuration
type Redis struct {
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	DB        int    `mapstructure:"db" json:"db" yaml:"db"`
	Pass      string `mapstructure:"pass" json:"pass" yaml:"pass"`
	Scheduler string `mapstructure:"scheduler" json:"scheduler" yaml:"scheduler"`
}
