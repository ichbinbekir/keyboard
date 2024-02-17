# ⌨️ Keyboard

Handle background keyboard inputs on windows for golang.

## ⚡️ Quickstart

```go
package main

import (
  "log"

  "github.com/ichbinbekir/keyboard"
)

const (
	_mouseLeftDown = 0x0002
	_mouseLeftUp   = 0x0004
)

func main() {
	kb := keyboard.New()

	defer kb.Close()

	kb.Handle(keyboard.NewKey('A'), func(state bool) {
		if state {
			keyboard.MouseEvent(_mouseLeftDown)
			keyboard.MouseEvent(_mouseLeftUp)
		}
	})

	ck := keyboard.NewKey('C')
	kb.Handle(keyboard.NewKey('B'), func(state bool) {
		if state {
			ck.Press()
		}
	})

	log.Println("listening keyboard")
	log.Fatal(kb.Listen())
}
```

## ⚙️ Installation

```bash
go get -u github.com/ichbinbekir/keyboard
```

## ⚠️ Warning

Key codes are missing in this library. You can get and use these codes from here:

<a href="https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes">Key Codes</a>,
<a href="https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mouse_event">Mouse Events</a>
