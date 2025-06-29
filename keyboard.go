package keyboard

type Keyboard struct {
	config Config
	Events chan Event
	closed bool
}

func New(configs ...Config) *Keyboard {
	kb := &Keyboard{config: mergeConfigs(configs...)}
	kb.Events = make(chan Event, kb.config.ChannelSize)
	go kb.readEvents()
	return kb
}

func (kb *Keyboard) Close() {
	kb.closed = true
	close(kb.Events)
}

func (kb Keyboard) Closed() bool {
	return kb.closed
}

func (kb Keyboard) GetConfig() Config {
	return kb.config
}

func GetState(key int) bool {
	return getState(key)
}

func SendInput(key int, state bool) {
	sendInput(key, state)
}
