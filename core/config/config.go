package config

// Config top level configuration
type Config struct {
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	MySQL     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Cassandra Cassandra `mapstructure:"cassandra" json:"cassandra" yaml:"cassandra"`
	Media     Media     `mapstructure:"media" json:"media" yaml:"media"`
}
