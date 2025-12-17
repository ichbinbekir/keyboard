package keyboard

type Keyboard struct {
	config    Config
	Events    chan Event
	Errors    chan error
	keyStates map[uint32]bool
	platform  platformSpecific
}

func New(cfg Config) *Keyboard {
	kb := &Keyboard{config: cfg}
	kb.Events = make(chan Event, kb.config.ChannelSize)
	kb.Errors = make(chan error)
	kb.keyStates = make(map[uint32]bool)
	go kb.readEvents()
	return kb
}

func (kb *Keyboard) Close() {
	close(kb.Events)
	close(kb.Errors)
}

func (kb *Keyboard) Config() Config {
	return kb.config
}

func KeyState(key int) bool {
	return getState(key)
}

func SendInput(key int, state bool) {
	sendInput(key, state)
}
