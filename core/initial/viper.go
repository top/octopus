package initial

import (
	"errors"
	"octopus/core/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() error {
	v := viper.New()
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config file not found")
		}
		return err
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		println("config file changed, reloading: " + e.Name)
		unmarshal(v)
	})
	return unmarshal(v)
}

func unmarshal(v *viper.Viper) error {
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		return err
	}
	return nil
}
