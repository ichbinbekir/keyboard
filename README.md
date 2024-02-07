# ⌨️ Keyboard

Handle background keyboard inputs on windows for golang.

## ⚡️ Quickstart

```go
package main

import (
  "log"

  "github.com/ichbinbekir/keyboard"
)

func main() {
	k := keyboard.NewKeyboard()

	defer k.Close()

	k.Handle(keyboard.NewKey('A'), func(state bool) {
		log.Println("A: ", state)
	})

	log.Fatal(k.Listen())
}

```

## ⚙️ Installation

```bash
go get -u github.com/ichbinbekir/keyboard
```
