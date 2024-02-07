package keyboard

type Key struct {
	Code  uintptr
	state bool
}

func NewKey(code uintptr) *Key {
	return &Key{Code: code}
}

func (k *Key) GetState() bool {
	return k.state
}

func (k *Key) Press() {
	_keybd_event.Call(k.Code)
}
