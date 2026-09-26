package main

import (
	"fmt"

	internal "bee/internal/bee"
)

func main() {
	bee := internal.NewBee()

	if err := bee.Run(); err != nil {
		panic(fmt.Errorf("error running bee: %w", err))
	}
}
