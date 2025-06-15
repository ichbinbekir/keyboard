package keyboard

import "reflect"

type Config struct {
	ChannelSize int
}

var defaultConfig = Config{
	ChannelSize: 8,
}

func mergeConfigs(cfgs ...Config) Config {
	config := defaultConfig
	if len(cfgs) == 0 {
		return config
	}

	configValue := reflect.ValueOf(&config).Elem()
	nField := configValue.NumField()
	for _, cfg := range cfgs {
		cfgValue := reflect.ValueOf(cfg)
		for field := range nField {
			value := cfgValue.Field(field)
			if !value.IsZero() {
				configValue.Field(field).Set(value)
			}
		}
	}

	return config
}
