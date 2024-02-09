package keyboard

import (
	"testing"
)

const (
	_mouseLeftDown = 0x0002
	_mouseLeftUp   = 0x0004
)

func TestMouse(t *testing.T) {
	if err := MouseEvent(_mouseLeftUp); err != nil {
		t.Error(err)
	}

	if err := MouseEvent(_mouseLeftDown); err != nil {
		t.Error(err)
	}
}
