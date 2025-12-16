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
