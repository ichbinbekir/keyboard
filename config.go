package keyboard

import "reflect"

type Config struct {
	ChannelSize          int
	SendRepeatedKeyDowns bool

	HandleKeyboard     bool
	HandleMouseButtons bool
	HandleMouseWheel   bool
	HandleMouseMove    bool
}

func DefaultConfig() Config {
	return Config{
		ChannelSize:    8,
		HandleKeyboard: true,
	}
}

// TODO: Not working
func mergeConfigs(cfgs ...Config) Config {
	config := DefaultConfig()
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
