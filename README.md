# ⌨️ Keyboard

Handle background keyboard inputs for golang.

## 🎯 Future

- [ ] Send input
- [ ] Handle mouse events
- [ ] Git Actions and tests
- [ ] Cross platform
- [ ] Key code enum

## ⚡️ Quickstart

```go
package main

import (
	"fmt"

	"github.com/ichbinbekir/keyboard"
)

func main() {
	cfg := keyboard.DefaultConfig()
	cfg.SendRepeatedKeyDowns = false
	cfg.HandleMouseButtons = true

	kb := keyboard.New(cfg)
	defer kb.Close()

	for {
		select {
		case event, ok := <-kb.Events:
			if !ok {
				break
			}
			fmt.Println(event)
		case err, ok := <-kb.Errors:
			if !ok {
				break
			}
			panic(err)
		}
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
