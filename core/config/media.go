package config

// Media level configuration
type Media struct {
	Facebook  Facebook  `mapstructure:"facebook" json:"facebook" yaml:"facebook"`
	Instagram Instagram `mapstructure:"instagram" json:"instagram" yaml:"instagram"`
}

// Instagram definition of Instagram declared above
type Instagram struct {
	ClientID     string `mapstructure:"client_id" json:"client_id" yaml:"client_id"`
	ClientSecret string `mapstructure:"client_secret" json:"client_secret" yaml:"client_secret"`
}

// Facebook definition of Facebook declared above
type Facebook struct {
	AppKey      string `mapstructure:"app_key" json:"app_key" yaml:"app_key"`
	AppSecret   string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
	RedirectURI string `mapstructure:"redirect_uri" json:"redirect_uri" yaml:"redirect_uri"`
}
