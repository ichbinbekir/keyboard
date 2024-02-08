package keyboard

import (
	"testing"
)

const (
	_mouseLeftDown = 0x0002
	_mouseLeftUp   = 0x0004
)

func TestKeyboard(t *testing.T) {
	kb := NewKeyboard()

	defer kb.Close()

	kb.Handle(NewKey('A'), func(state bool) {
		if state {
			MouseEvent(_mouseLeftDown)
			MouseEvent(_mouseLeftUp)
		}
	})

	ck := NewKey('C')
	kb.Handle(NewKey('B'), func(state bool) {
		if state {
			ck.Press()
		}
	})

	t.Log("listening keyboard")
	t.Fatal(kb.Listen())
}
