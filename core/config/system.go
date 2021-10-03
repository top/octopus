package config

// System level configuration
type System struct {
	Env       string    `mapstructure:"env" json:"env" yaml:"env"`
	TLS       bool      `mapstructure:"tls" json:"tls" yaml:"tls"`
	CertFile  string    `mapstructure:"cer" json:"cer" yaml:"cer"`
	KeyFile   string    `mapstructure:"key" json:"key" yaml:"key"`
	Addr      string    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType    string    `mapstructure:"dbType" json:"dbType" yaml:"dbType"`
	Transform Transform `mapstructure:"transform" json:"transform" yaml:"transform"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
}
