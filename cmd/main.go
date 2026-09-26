package main

import (
	"fmt"
	"os"

	internal "bee/internal/bee"
)

func main() {
	bee := internal.NewBee()

	if err := bee.Run(); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
