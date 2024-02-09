package keyboard

import (
	"testing"
)

func TestKey(t *testing.T) {
	ak := NewKey('A')

	if err := ak.Press(); err != nil {
		t.Error(err)
	}
}
