package main

import (
	"flag"
	"fmt"

	internal "bee/internal/bee"
)

func main() {
	start := flag.String("start", "./", "start a new bee directory.")
	flag.Parse()

	bee := internal.NewBee()
	if err := bee.Run(); err != nil {
		fmt.Errorf("error running bee: %w", err)
	}

	fmt.Printf("start value: %s", *start)
}
