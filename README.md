# ⌨️ Keyboard

Handle background keyboard inputs on windows for golang.

## 🕦 Future

- [ ] Send click event
- [ ] Handle mouse events

## ⚡️ Quickstart

```go
package main

import (
  "log"

  "github.com/ichbinbekir/keyboard"
)

func main() {
	kb := New(/* Config{} */)
	defer kb.Close()

	for event := range kb.Events {
		log.Println(event)
	}
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
