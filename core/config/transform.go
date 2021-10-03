package config

// Transform level configuration
type Transform struct {
	Queue    string `mapstructure:"queue" json:"queue" yaml:"queue"`
	Capacity int    `mapstructure:"capacity" json:"capacity" yaml:"capacity"`
}
