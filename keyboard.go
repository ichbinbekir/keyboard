package keyboard

type Handler func(bool)

type Keyboard struct {
	bindings  map[*Key]Handler
	listining bool
}

func NewKeyboard() *Keyboard {
	return &Keyboard{bindings: make(map[*Key]Handler)}
}

func (kb *Keyboard) Close() {
	kb.listining = false
}

func (kb *Keyboard) Handle(k *Key, h Handler) {
	kb.bindings[k] = h
}

func (kb *Keyboard) IsListining() bool {
	return kb.listining
}

func (kb *Keyboard) Listen() error {
	kb.listining = true

	for kb.listining {
		for key, handler := range kb.bindings {
			state, err := getKeyState(key.Code)
			if err != nil {
				return err
			}

			if state != key.state {
				key.state = state

				if handler != nil {
					handler(state)
				}
			}
		}
	}

	return nil
}
