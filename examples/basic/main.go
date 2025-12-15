package main

import (
	"fmt"

	"github.com/ichbinbekir/keyboard"
)

func main() {
	kb := keyboard.New(keyboard.Config{HandleMouseButtons: true})
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
