package keyboard

import (
	"testing"
	"time"
)

func TestKeyboard(t *testing.T) {
	kb := New()

	kb.Handle(NewKey('A'))

	if err := kb.Listen(func() { time.Sleep(time.Second); kb.Close() }); err != nil {
		t.Error(err)
	}
}
