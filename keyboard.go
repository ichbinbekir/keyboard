package keyboard

type Handler func(bool)

type Keyboard struct {
	bindings  map[*Key]Handler
	listining bool
}

func New() *Keyboard {
	return &Keyboard{bindings: make(map[*Key]Handler)}
}

func (kb *Keyboard) Close() {
	kb.listining = false
}

func (kb *Keyboard) Handle(k *Key, handlers ...Handler) {
	if len(handlers) == 0 {
		kb.bindings[k] = nil
	}

	for _, handler := range handlers {
		kb.bindings[k] = handler
	}
}

func (kb *Keyboard) IsListining() bool {
	return kb.listining
}

func (kb *Keyboard) Listen(onLoads ...func()) error {
	kb.listining = true

	for _, onLoad := range onLoads {
		go onLoad()
	}

	for kb.listining {
		for key, handler := range kb.bindings {
			state, err := getKeyState(key.Code)
			if err != nil {
				return err
			}

			if state != key.state {
				key.state = state

				if handler != nil {
					go handler(state)
				}
			}
		}
	}

	return nil
}
