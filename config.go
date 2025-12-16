package keyboard

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
		ChannelSize:          8,
		SendRepeatedKeyDowns: true,
		HandleKeyboard:       true,
	}
}
